package main

import "fmt"

func main() {

	i:=[]int{10,50,60,70,80,90}

	b:= i[2:6]

	b[0] = 200

	fmt.Println("i",i)
	fmt.Println("b",b)


}
