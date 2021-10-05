package main

import "fmt"

type students [10]string

func main() {

	var a [10]string = [10]string{}
	var name students
	
	b := students{"abc", "xyz"}
	_ = b

	if a == name { /// comparison only possible if types are same
		fmt.Println("same")
	}

}
