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
	cacheRepo "mivar_robot_api/internal/repo/cache"
	manager "mivar_robot_api/internal/service/model_manager"
	"mivar_robot_api/pkg/cache"
	"mivar_robot_api/pkg/generator"
)

const (
	wait        = time.Second * 5
	initTimeout = time.Second * 15
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		configPath = "config.yaml"
	}
	conf, err := configer.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("Cant laod config: %v", err))
	}

	logger := logrus.New()
	wimiCli, err := mivar.New(
		mivar.ClientConfig{
			BaseURL:    "http://127.0.0.1:8092",
			Timeout:    time.Second * 10,
			HTTPClient: &http.Client{},
		})
	modelGenerator := generator.NewGenerator()
	inMemCache := cache.NewCache()
	cacheRepository := cacheRepo.New(inMemCache)

	modelManager := manager.New(logger, cacheRepository, modelGenerator, wimiCli, *conf)

	ctx, cancel := context.WithTimeout(context.Background(), initTimeout)
	defer cancel()

	err = modelManager.LoadModels(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to init: %v", err))
	}

	runService(logger)

}

func runService(log *logrus.Logger) {
	calcPathHandler := calc_path.NewCalcPathHandler(log)
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/calcPath", calcPathHandler.Handle).Methods("POST")

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
