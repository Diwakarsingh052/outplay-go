package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int) //unbuffered channel
	wg.Add(4)
	go add(2, 2, c)
	go sub(4, 2, c)
	go mult(5, 5, c)
	go calc(c)
	wg.Wait()

}

func add(a, b int, c chan int) {
	defer wg.Done()
	sum := a + b
	c <- sum //send
}

func sub(a, b int, c chan int) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	sum := a - b
	c <- sum //send
}
func mult(a, b int, c chan int) {
	defer wg.Done()
	sum := a * b
	c <- sum //send
}

func calc(c chan int) {
	defer wg.Done()
	fmt.Println("In calc")
	x, y, z := <-c, <-c, <-c // blocking ops // receive
	fmt.Println(x + y + z)
}
