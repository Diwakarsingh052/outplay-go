package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	const cap int = 2
	const work = 10

	ch := make(chan string, work) // buffered channel
	sem := make(chan bool, cap)

	for e := 1; e <= work; e++ {

		go func(e int) { // g10 ,g20 , g30
			sem <- true // [t,t] //blocking ops
			{
				time.Sleep(2 * time.Second)
				ch <- "task" + strconv.Itoa(e)
			}
			<-sem // receive  // pull out the values

		}(e)
	}

	for i := 1; i <= work; i++ {

		p := <-ch
		fmt.Println("recv", p)
	}

}
