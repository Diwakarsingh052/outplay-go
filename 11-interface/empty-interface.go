package main

import "fmt"

func main()  {

	var i interface{}
	i = 2

	i = true
	i = "abc"
	//var check bool
	a, ok := i.(int)

	if !ok {
		fmt.Println(a)
	}

	fmt.Println(i)
	fmt.Printf("%T",i)

}
