package main

import "fmt"

func main() {

	show("", 10, 20, 50, 60)
}

func show(s string, a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
