package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string, 1) //buffered channel
	wg.Add(1)
	go func() {
		ch <- "Adding work here"
		wg.Done()
	}()

	//wg.Add(1) not added
	go func() {
		time.Sleep(5 * time.Second)
		<-ch
		time.Sleep(time.Second)
		fmt.Println("value received", <-ch)
	}()
	wg.Wait()
}

// A-> Deadlock
// B-> wait 5 second and exits
// c-> everything is fine, main will not wait and exit
// d-> wait 6 seconds and print value received
