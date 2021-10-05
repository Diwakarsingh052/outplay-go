package main

import "fmt"

func main() {

	i := []int{10, 20, 30}

	fmt.Println(len(i), cap(i))
	fmt.Printf("address %p\n", i)

	i = append(i, 40)
	fmt.Printf("address %p\n", i)
	fmt.Println(len(i), cap(i))

	i = append(i, 400,600,700)
	fmt.Printf("address %p\n", i)
	fmt.Println(len(i), cap(i))
	fmt.Println(i)

}
