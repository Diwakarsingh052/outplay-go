package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {

		c1 <- "one"
	}()

	go func() {
		//time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		c3 <- "three"
	}()

	for i:=1 ; i <=3 ; i++ {


	select {
	case x := <-c1:
		fmt.Println(x)
	case y := <-c2:
		fmt.Println(y)
	case z := <-c3:
		fmt.Println(z)
	}

	}

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)

}
