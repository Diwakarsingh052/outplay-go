package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	user := map[string]string{"first_name": "Raj"}
	u, _ := json.Marshal(user)

	request, err := http.NewRequest(http.MethodPost, `http://httpbin.org/post`, bytes.NewReader(u))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()

	request = request.WithContext(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("Content-Type", "application/json")

	//client := http.Client{}
	//client.Do()
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	data, _ := io.ReadAll(resp.Body)

	fmt.Println(string(data))

}
