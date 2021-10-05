package main

import "fmt"

func main() {

	var a [5]int //fixed size // array size cannot grow // size cannot be changed at runtime
	fmt.Println(a)
	a[0] = 100

	b := [10]int{10, 100, 200, 300}
	fmt.Println(b)

	var c = [3]int{2, 6, 8}

	//c[10] = 200
	fmt.Println(c)

}
