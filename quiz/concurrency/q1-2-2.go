package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch2 := make(chan string)

	ch2 <- "doing work" //blocked // avoid send in main go when using unbuffered channel
	wg.Add(1)
	go func() {

		fmt.Println(<-ch2)
		wg.Done()
	}()
	wg.Wait()
}

// a) Deadlock
// b) doing work output
