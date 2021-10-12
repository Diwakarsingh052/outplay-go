package main

import "fmt"

type Walk interface {
	Walk()
}

type Run interface {
	Run()
}

type WalkRunner interface {
	Walk
	Run
}

type Human struct{}

func (h Human) Walk() {
	fmt.Println("walker")
}
func (h Human) Run() {
	fmt.Println("runner")
}
func (robo) Walk() {
	fmt.Println("robo walker")
}
func (robo) Run() {
	fmt.Println("robo runner")
}

type robo struct{}

func main() {
	h := Human{}
	_ = h
	ro := robo{}
	var r Run
	var w Walk
	var wr WalkRunner

	wr = ro
	r = wr // assign walk runner to run
	r.Run()
	w = h
	//wr = w

	hu, ok := w.(robo)

	if !ok {
		fmt.Println("not human")
		return

	}

	wr = hu
	wr.Run()
	wr.Walk()

}

//func abc(r Run) {
//
//}
