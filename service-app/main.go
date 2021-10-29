package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/auth"
	"service-app/data/user"
	"service-app/database"
	"service-app/handlers"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "Users : ", log.LstdFlags)
	startApp(l)
}

func startApp(log *log.Logger) error {

	// =========================================================================
	// Initialize authentication support
	log.Println("main : Started : Initializing authentication support")

	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	a, err := auth.NewAuth(privateKey, "RS256")
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}
	// =========================================================================
	// Start Database

	log.Println("main: Initializing database support")
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}
	c := database.Config{
		Host:     os.Getenv("Host"),
		User:     os.Getenv("Name"),
		Password: os.Getenv("Password"),
		Port:     os.Getenv("Port"),
		DBName:   os.Getenv("database"),
		TLSMode:  os.Getenv("TLSMode"),
	}
	db, err := database.Open(c)

	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}

	uDB := user.NewDbService(db)
	defer func() {
		log.Printf("main: Database Stopping : ")
		db.Close()
	}()

	shutdown := make(chan os.Signal, 1) // buffered channel
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:         ":8080",
		Handler:      handlers.API(shutdown, log, a, uDB), // registering handlers functions  log, a, db, uDB
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
