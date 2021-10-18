package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ch := make(chan string, 1)

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		ch <- "work"

	}()

	select {
	case p := <-ch:
		fmt.Println(p)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	fmt.Println("hello")
	wg.Wait()

}
