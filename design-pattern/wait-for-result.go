package main

import (
	"fmt"
	"time"
)

func main() {

	waitForResult()


}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(3* time.Second)
		ch <- "result"

	}()

	r:= <-ch // unknown latency

	fmt.Println(r)


}