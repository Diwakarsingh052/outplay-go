package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string, 1) // buffered channel
	wg.Add(1)
	go func() {
		ch <- "Adding work here"
		fmt.Println("receiving here", <-ch)
		wg.Done()
	}()

	wg.Wait()
}

// A-> Deadlock
// B-> everything is fine
// c-> Syntax error
