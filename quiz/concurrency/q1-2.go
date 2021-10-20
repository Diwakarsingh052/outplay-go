package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string)

	wg.Add(1)
	go func() {
		ch1 <- "Adding work here" // block
		fmt.Println("receiving here", <-ch1)
		wg.Done()
	}()


	wg.Wait()
}

// A-> Deadlock
// B-> everything is fine
// c-> Syntax error
