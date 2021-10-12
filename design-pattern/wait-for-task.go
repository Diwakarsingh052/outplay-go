package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	waitForTask()
}

type resp struct {
	url string
	res *http.Response
	err error
}

func waitForTask() {
	ch := make(chan string)
	respCh := make(chan resp)

	go func() {
		t := <-ch // this is going to wait until it doesn't receive anything in channel
		//var err error
		res, err := http.Get(t)
		a := resp{
			url: t,
			res: res,
			err: err,
		}
		respCh <- a

	}()

	time.Sleep(2 * time.Second)
	ch <- "https://pkg.go.dev/"

	respStruct := <-respCh
	if respStruct.err != nil {
		log.Println(respStruct.err)
		return
	}
	b,err:= io.ReadAll(respStruct.res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

}
