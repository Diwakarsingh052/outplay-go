package sum

import (
	"fmt"
)

//a:="hello" // no shorthand in global scope

var apiKey string = "qwertyuio"

func Add() { // making func first letter as uppercase exports it
	fmt.Println("I am calling add")
	sub() // you can call unexported func in the same directory

	fmt.Println(apiKey)
	//os.Args
}
