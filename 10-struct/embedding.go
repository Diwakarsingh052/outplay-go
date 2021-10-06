package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	user // not embedding
	role []string
}

func (u user) notify() {
	fmt.Println("I am from user")
}

func (a admin) notify() {
	fmt.Println("I am from admin")
}

func main() {

	a := admin{
		user: user{
			name:  "dev",
			email: "dev@email.com",
		},
		role: []string{"update"},
	}

	a.user.notify()
	a.notify()

}
