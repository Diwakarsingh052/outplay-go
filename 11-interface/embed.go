package main

type DB interface {
	read() string
}

type users struct {
	name string
	//DB //nil
}

func (u *users) read() string {
	return ""
}

func main() {
	u := users{}
	u.read()
}

func NewUser(name string) *users {
	u:= users{name: "abc"}

	return &u
}
