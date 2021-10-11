package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("Sending a notification to %s %s %s", u.name, u.email, string(p))
	return len(p), nil
}

func main() {
	//f,err:= os.OpenFile("abc",0,0)
	u:= user{
		name:  "abc",
		email: "abc@email.com",
	}
	l:= log.New(u,"",0)
	l.Println("my data")
}
