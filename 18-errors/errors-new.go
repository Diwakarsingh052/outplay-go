package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFolderNotFound = errors.New("file not found in root folder")

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
		return nil, fmt.Errorf("%v :: %v", err, ErrFolderNotFound) // deprecated
	}
	return f, nil
}
