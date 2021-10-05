package main

import "fmt"

func main() {

	day := "mon"

	switch day {

	case "mon":
		fmt.Println("monday")
		fallthrough
	case "tues":
		fmt.Println("tuesday")
		fallthrough
	default:
		fmt.Println("nothing matches")

	}

}
