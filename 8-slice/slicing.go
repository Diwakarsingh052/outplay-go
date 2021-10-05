package main

import "fmt"

func main()  {

	a:= []int{10,70,60,90,100,200}
	b:= a[:] // take whole slice
	b= a[2:]
	b = a[:3]

	b = a[2:4] // index:length
	fmt.Println(b)




}
