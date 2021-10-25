package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrFolderNotFound = errors.New("file not found in root folder")

func main() {
	_, err := openFile("xyz.txt")
	//if err != nil {
	//	log.Println(err)
	//	os.Exit(1)
	//}
	if errors.Is(err, ErrFolderNotFound) {
		fmt.Println("yes its their")
	} else {
		fmt.Println("no its not")
	}
	var err error
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	//wrapedError := fmt.Errorf("just checking on errors")

	if err != nil {
		return nil, fmt.Errorf("%v :: %w", err, ErrFolderNotFound) // %w is wrapping of the error // only one error could be wrapped at a time
	}
	return f, nil
}
