package main

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = errors.New("database.GetUser : record was not found with the email address provided")
var user map[int]string = make(map[int]string)

func main() {

	//var err error
	//_, err := strconv.Atoi("A")
	//fmt.Printf("%v", err.Error())

	_, err := FetchRecord(1)

	fmt.Println(err)

}

func FetchRecord(id int) (string, error) {
	u, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	return u, nil

}
