package main

import (
	"encoding/json"
	"log"
	"os"
)

type Perms map[string]bool
type person struct {
	FirstName string `json:"first_name"`
	Password  string `json:"-"`
	Perms     `json:"perms,omitempty"`
}

func main() {

	users := []person{

		{
			FirstName: "Roy",
			Password:  "abc",
			Perms:     Perms{"admin": true},
		},

		{
			FirstName: "Raj",
			Password:  "qwe",
			Perms:     Perms{"write": false},
		},

		{
			FirstName: "Pulkit",
			Password:  "rty",
		},
	}

	byte, err := json.Marshal(users)

	if err != nil {
		return
	}

	f, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)
	//fmt.Println(string(byte))
	if err != nil {
		log.Fatalln(err)
	}

	f.Write(byte)
}
