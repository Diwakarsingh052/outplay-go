package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//fmt.Println(runtime.GOMAXPROCS(6))
	//fmt.Println(runtime.GOMAXPROCS(6))
	wg.Add(20)
	for i := 1; i <= 10; i++ {
		//wg.Add(1)
		go result(i) //
	}
	for i := 1; i <= 10; i++ {
		//wg.Add(1)
		go result(i) //
	}
	wg.Wait() // until counter is not 0 wait here
}

func result(i int) {
	defer wg.Done() // decrease the counter
	fmt.Println("hello ", i)

}
