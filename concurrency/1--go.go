package main

import (
	"fmt"
	"time"
)

func main()  {

	go hello()
	time.Sleep(time.Second) //bad idea
}

func hello()  {
	fmt.Println("hello")
}
