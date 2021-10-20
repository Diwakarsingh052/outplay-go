package main

import (
	"log"
	"os"
)

func main() {
	_, err := openFile("xyz.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	return f, nil
}
