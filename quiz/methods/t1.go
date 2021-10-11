package main

import (
	"fmt"
	"os"
)

type files struct {
	fs *os.File
}

func (f files) open(name string) error {
	var err error

	f.fs, err = os.OpenFile(name, os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	return nil

}

func (f files) Create(data string) error {

	fmt.Println("fs", f.fs)
	_, err := f.fs.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}



// options
/*
a-> Creating err , Writing nil
b -> Creating nil, Writing nil  //no error in creating file and writing to it
c -> Creating nil, Writing err
d-> Creating err, Writing err
*/


func main() {
	var f files //nil
	fmt.Println("creating", f.open("abc.txt"))
	fmt.Println("writing", f.Create("hello"))

}
