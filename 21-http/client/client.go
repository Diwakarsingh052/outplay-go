package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {


	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	//http://localhost:8080
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	req = req.WithContext(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}
	io.Copy(os.Stdout, res.Body)

}
