package main

import "fmt"

func main()  {
	a:=19
	update(&a)
	fmt.Println(a)
}

func update(a *int)  {
	*a= 100
	//a++

}
