package main

import "fmt"

func main() {

	CheckType("20")

}

func CheckType(a interface{}) {
	switch v := a.(type) {

	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float", v)
	default:
		fmt.Println("don't know", v)

	}

}
