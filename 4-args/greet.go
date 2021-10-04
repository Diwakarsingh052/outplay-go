package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	details := os.Args[1:]

	if len(details) < 2 {
		log.Println("please provide your name and age")
		return // it will stop exec of current function
	}
	name := details[0]
	age := details[1]
	fmt.Println(name, age)

}
