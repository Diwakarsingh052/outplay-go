package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", Home) // registers the pattern and the functions

	log.Fatalln(http.ListenAndServe(":8080", nil))

}

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello this is the home page")

}
