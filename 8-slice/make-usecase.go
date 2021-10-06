package main

import "fmt"

func main() {

	var data []int = make([]int, 0, 100000)

	lastCap := cap(data)
	var allocation int
	for r := 0; r < 100000; r++ {
		//data[r] = r
			data = append(data, r)

		if lastCap != cap(data) {
			allocation++

			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Add [%p] Cap [%d - %v]\n", data, cap(data), capCh)

		}

	}
	fmt.Println(allocation)

}
