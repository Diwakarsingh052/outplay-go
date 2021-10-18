package main

import "fmt"

func main() {
	var names map[int]int // nil
	names = make(map[int]int)
	names[1] = 10

	a,ok := names[1]

	if !ok {
		fmt.Println("key not found")
		return
	}
	fmt.Println(a)

	fmt.Println(names)
	//fmt.Printf("%#v",names)

}
