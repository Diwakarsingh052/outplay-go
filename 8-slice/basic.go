package main

import "fmt"

func main() {

	var i []int // reference // nil // length

	//i[1] = 100
	fmt.Println(i)
	fmt.Printf("%#v\n",i)

	if i == nil {
		fmt.Println("it is nil")
	}

}
