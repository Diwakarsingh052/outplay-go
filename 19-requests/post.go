package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	user := map[string]string{"first_name": "Raj"}
	u, _ := json.Marshal(user)

	request, err := http.NewRequest(http.MethodPost, `http://httpbin.org/post`, bytes.NewReader(u))

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
