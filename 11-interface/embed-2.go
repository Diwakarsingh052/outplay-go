package main

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

}
func (h Human) Run() {

}

func main() {
	h := Human{}
	abc(h)
}

func abc(r WalkRunner) {

}
