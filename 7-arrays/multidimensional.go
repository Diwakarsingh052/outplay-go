package main

import "fmt"

func main() {

	a := [2][5]int{
		{10, 100},
		{200},
	}

	for _, v := range a {
		fmt.Println(v)

	}

}
