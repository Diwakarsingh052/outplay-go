package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//defer fmt.Println("1")
	defer fmt.Println("2")
	//
	//fmt.Println("3")
	//fmt.Println("4")
	f, err := os.OpenFile("abcd", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

}
