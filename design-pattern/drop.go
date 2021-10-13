package main

import (
	"fmt"
	"sync"
)

// https://golang.org/ref/spec#Send_statements
//An Unbuffered channel gives you a Guarantee that a signal being sent has been received.
//Because the Receive of the signal Happens Before the Send of the signal completes.
//

////No Guarantee
//A Buffered channel of size >1 gives you No Guarantee that a signal being sent has been received.
//Because the Send of the signal Happens Before the Receive of the signal completes.

func main() {
	var wg sync.WaitGroup

	const cap = 2
	ch := make(chan string,cap) // buffered so send happens before receive
	wg.Add(1)
	go func() { //ready
	defer wg.Done()
		for p := range ch { // you can recv values over a closed channel but you cannot send values over a closed channel
			fmt.Println("recv ", p)
		}

	}()

	for w := 1; w <= 10; w++ {

		select {
		case ch <- "task":
			fmt.Println("sent a signal")
		default:
			fmt.Println("Drop") //if unbuffered channel is used it will always drop because receive happens first and select will pick default case because sender is not ready yet

		}

	}
	close(ch) // close channels from where you are sending the data
	wg.Wait()

}
