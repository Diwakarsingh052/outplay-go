package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://loripsum.net/api")

	if err != nil {
		log.Println(err)
		return
	}
	data, _ := io.ReadAll(resp.Body)

	fmt.Println(string(data))

}
