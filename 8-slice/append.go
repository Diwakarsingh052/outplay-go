package main

import "fmt"

func main() {

	a := []int{10}
	b := []int{16, 80, 50}

	a = append(a, 40, 50)
	a = append(a, b...)

	fmt.Println(a)

}
