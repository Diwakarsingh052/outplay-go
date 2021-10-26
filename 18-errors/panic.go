package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	abc()

	fmt.Println("fine")

}

func abc() {
	defer recoverFileNotFound()
	panic(" i have a panic")
}

func recoverFileNotFound() {
	r := recover()
	if r != nil {
		fmt.Println("recovered", r)
		fmt.Printf("%s", debug.Stack())
	}

}
