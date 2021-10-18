package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("emp is working", v)
		}

	}()

	go func() {
		defer wg.Done()
		for w := 1; w <= 10; w++ {

			ch <- "work" + strconv.Itoa(w)
		}
		//close(ch) //close channel where we are sending the data

	}()

	wg.Wait()

}
