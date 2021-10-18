package main

import (
	"fmt"
	"sync"
	"time"
)

var cabs = 1
var wg sync.WaitGroup

func main() {
	m := &sync.Mutex{} //it protects when you are doing write ops
	names := []string{"a", "b", "c", "d"}
	wg.Add(4)
	for _, n := range names {
		go bookCab(n, m)
	}
	wg.Wait()
}

func bookCab(name string, m *sync.Mutex) {
	defer wg.Done()

	fmt.Println("welcome to the website", name)
	m.Lock()
	if cabs >= 1 {
		fmt.Println("cab is available for", name)
		time.Sleep(time.Second)
		fmt.Println("Booking confirmed", name)
		cabs--
	} else {
		fmt.Println("cab is not available for", name)
	}
	m.Unlock()

}
