package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string)
	ch2 := make(chan string)
	wg.Add(1)
	go func() {
		fmt.Println("receiving here", <-ch1) // forgotten receiver
		ch2 <- "Adding work here"
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch2
	}()
	wg.Wait()
}

// A-> Deadlock
// B-> everything is fine
// c-> Syntax error







