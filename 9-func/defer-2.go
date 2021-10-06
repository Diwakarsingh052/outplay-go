package main

import "fmt"

func main() {

	abc()
	defer fmt.Println("1 from defer")
	return

	defer fmt.Println("2 from defer")

	fmt.Println("Hello")

}

func abc() {
	defer fmt.Println("10")
}
