package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fanOut()
}

func fanOut()  {

	emps := 10

	ch := make(chan string,emps)

	for work:= 1; work<=emps ; work++ {
		go func(w int) { // ready state
			ch <- "work " + strconv.Itoa(w)
		}(work)
	}

	for emps > 0 {
		fmt.Println(<-ch)
		emps--
	}





}
