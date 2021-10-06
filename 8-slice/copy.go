package main

import "fmt"

func main() {

	i := []int{10, 20, 30, 40, 50}

	b := make([]int, len(i))

	copy(b, i[1:4])

	b[0] = 900

	fmt.Println(i)
	fmt.Println(b)

}
