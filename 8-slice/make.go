package main

import "fmt"

func main() {

	var i []int
	i = make([]int, 0, 10)
	inspectSlice("i", i)

	i = append(i, 100)
	inspectSlice("i", i)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))

	for i := range slice {
		fmt.Printf(" [%d] %p %v", i, &slice[i], slice[i])
	}

	fmt.Println()
	fmt.Println(slice)
	fmt.Println()

}
