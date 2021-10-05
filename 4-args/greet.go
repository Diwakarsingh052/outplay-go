package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	details := os.Args[1:]

	if len(details) < 3 {
		log.Println("please provide your name and age marks")
		return // it will stop exec of current function
	}
	name := details[0]
	ageString := details[1]
	marksString := details[2]
	//fmt.Println(name, age)

	var err error //nil  // nil -> no err
	fmt.Println(err)
	age, err := strconv.Atoi(ageString) //err -> invalid syntax

	if err != nil {
		log.Println(err)
		fmt.Println(age)
		return
	}

	marks, err := strconv.Atoi(marksString) // nil

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(name, age, marks)

}
