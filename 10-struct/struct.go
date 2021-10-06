package main

import "fmt"

type user struct {
	name  string //fields
	marks []int
}

type student struct{}

func main() {

	var u user = user{}

	fmt.Println(u)
	//s := user{
	//	name:  "abc",
	//	marks: []int{100, 700, 600},
	//}
	//fmt.Println(s)

	u.name = "Raj"
	u.show()
	u.update()
	u.show()
	s := student{}

	s.update()
}

func (u user) show() { // receiver
	fmt.Println(u)
}

func (u *user) update() { // receiver
	u.name = "abc"
}

func (s *student) update() {

}
