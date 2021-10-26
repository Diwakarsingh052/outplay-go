package main

import (
	"context"
	"fmt"
	"http-go/mylog"
	"log"
	"net/http"
	"time"
)

func main() {
	m := mylog.Middleware(randomData)
	http.HandleFunc("/", m)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func randomData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int64(66), int64(999))
	mylog.Println(ctx, "started")
	defer mylog.Println(ctx, "completed")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Random Data")
	case <-ctx.Done():
		err := ctx.Err()
		mylog.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
