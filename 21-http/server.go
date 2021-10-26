package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", randomData)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func randomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("started")
	defer log.Println("completed")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Random Data")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
