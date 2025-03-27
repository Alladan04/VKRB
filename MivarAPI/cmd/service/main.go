package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"mivar_robot_api/internal/controller/http/calc_path"
)

func main() {
	var wait = time.Second * 5

	calcPathHandler := calc_path.NewCalcPathHandler()
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
