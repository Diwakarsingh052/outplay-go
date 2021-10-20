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
		ch <- "Adding work here 1st time"
		ch <- "Adding work here 2nd time"
		fmt.Println("receiving here", <-ch)
		wg.Done()
	}()

	wg.Wait()
}

// A-> Deadlock due to line 13
// B-> everything is fine
// c-> Deadlock due to line 14
