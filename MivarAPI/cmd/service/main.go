package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/client/mivar"
	"mivar_robot_api/internal/config"
	"mivar_robot_api/internal/controller/http/calc_path"
	restoreHandler "mivar_robot_api/internal/controller/http/restore"
	updateMap "mivar_robot_api/internal/controller/http/update_map"

	cacheRepo "mivar_robot_api/internal/repo/cache"
	"mivar_robot_api/internal/repo/persistent"
	manager "mivar_robot_api/internal/service/model_manager"
	calculate "mivar_robot_api/internal/usecase/calculate_path"
	"mivar_robot_api/internal/usecase/restore"
	updateMapUc "mivar_robot_api/internal/usecase/update_map"
	"mivar_robot_api/pkg/cache"
	"mivar_robot_api/pkg/generator"
)

const (
	wait        = time.Second * 3
	initTimeout = time.Second * 15
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	logger := logrus.New()

	if configPath == "" {
		configPath = "config.yaml"
	}

	logger.Info("config path: ", configPath)
	conf, err := configer.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("Cant laod config: %v", err))
	}

	logger.Info("loaded config")
	wimiCli, err := mivar.New(
		mivar.ClientConfig{
			BaseURL:    "http://wimi-service:8092",
			Timeout:    time.Second * 10,
			HTTPClient: &http.Client{},
		})
	modelGenerator := generator.NewGenerator()
	inMemCache := cache.NewCache()
	inMemCacheLabirint := cache.NewCache()
	cacheRepository := cacheRepo.New(inMemCache, inMemCacheLabirint)

	modelManager := manager.New(logger, cacheRepository, modelGenerator, wimiCli, *conf)

	ctx, cancel := context.WithTimeout(context.Background(), initTimeout)
	defer cancel()

	err = modelManager.LoadModels(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to init: %v", err))
	}

	runService(logger, cacheRepository, modelGenerator, wimiCli, modelManager, *conf)

}

func runService(log *logrus.Logger, inMemCache *cacheRepo.Repo, modelGenerator *generator.Generator, wimiCli *mivar.Client, manager *manager.Manager, conf configer.Config) {
	calculateUsecase := calculate.New(log, inMemCache, modelGenerator, wimiCli, manager)
	calcPathHandler := calc_path.NewCalcPathHandler(log, calculateUsecase)

	updateMapUsecase := updateMapUc.New(log, inMemCache, modelGenerator, wimiCli)
	updateMapHandler := updateMap.New(log, updateMapUsecase)

	persistentRepo := persistent.New(conf)
	restoreUc := restore.New(log, inMemCache, persistentRepo, wimiCli)
	restoreHandler := restoreHandler.New(log, restoreUc)

	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/calc_path", calcPathHandler.Handle).Methods("POST")
	r.HandleFunc("/update_map", updateMapHandler.Handle).Methods("POST")
	r.HandleFunc("/restore_map", restoreHandler.Handle).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 10,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()
	// Block until the shutdown is complete or the context deadline is reached.
	<-ctx.Done()
	log.Println("shutting down")
	os.Exit(0)

}
