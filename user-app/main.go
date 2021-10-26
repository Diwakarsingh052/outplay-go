package main

import (
	"net/http"
	"user-app/handlers"
)

func main() {

	http.HandleFunc("/user", handlers.GetUser)
	http.ListenAndServe(":8080", nil)
}
