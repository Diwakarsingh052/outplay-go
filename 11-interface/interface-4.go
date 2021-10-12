package main

type abc interface {
	read()
}

type person struct {
	name string
}

func (person) read() {

}
func (person) Put(string) error {
	return nil
}

func (person) write() {

}

func main() {

	var in abc
	var p person

	in = p
	_ = in

	//in.

}
