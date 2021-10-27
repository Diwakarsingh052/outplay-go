package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/handlers"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "Users : ", log.LstdFlags)
	startApp(l)
}

func startApp(log *log.Logger) error {
	shutdown := make(chan os.Signal, 1) // buffered channel
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:         ":8080",
		Handler:      handlers.API(shutdown, log), // registering handlers functions  log, a, db, uDB
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 8000 * time.Second,
	}

	serverErrors := make(chan error, 1)
	// Start the service listening for requests.
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()

	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel() // cleaning up the resources of timers

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil { // first trying to cleanly shutdown
			api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}
	}
	return nil

}
