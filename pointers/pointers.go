package main

import "fmt"

func main() {

	var i int
	var p *int
	fmt.Println(p)
	p = &i
	*p = 100
	fmt.Println(p)
	fmt.Println(&i)
	fmt.Println(i)

}
