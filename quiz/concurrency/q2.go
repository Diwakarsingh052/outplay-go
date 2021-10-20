package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1 // forgotten sender
	}()
	fmt.Println("hello")
	wg.Wait()
}

// a) hello as output and deadlock
// b) deadlock only
// c) hello as output and no deadlock
// d) hello output only
