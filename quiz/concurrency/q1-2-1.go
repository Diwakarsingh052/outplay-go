package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch2 := make(chan string)

	wg.Add(1)
	go func() {

		fmt.Println(<-ch2)
		wg.Done()
	}()

	ch2 <- "doing work"
	wg.Wait()
}

// a) Deadlock
// b) doing work output
