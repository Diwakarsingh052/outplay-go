package main

import "fmt"

func main() {
	i := []int64{54, 67, 89, 600, 800, 500}
	b := i[1:4:4]
	b[0] = 999 // this will still affect the parent slice or backing array
	b = append(b, 300)


	inspectSlice("i", i)
	inspectSlice("b", b)
}

func inspectSlice(name string, slice []int64) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))

	for i := range slice {
		fmt.Printf(" [%d] %p %v", i, &slice[i], slice[i])
	}

	fmt.Println()
	fmt.Println(slice)
	fmt.Println()

}
