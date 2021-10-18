package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type cab struct {
	driver int
	rw     sync.RWMutex
	wg     sync.WaitGroup
}

func (c *cab) getCabDrivers() {
	defer c.wg.Done()

	c.rw.RLock()         // when a goroutine is reading no other cad do a write ops // this is read lock
	defer c.rw.RUnlock() // after unlocking anyone can write
	time.Sleep(time.Second)
	fmt.Println("driver", c.driver)

}

func (c *cab) bookCab(name string) {
	defer c.wg.Done()
	c.rw.Lock() // locking for writes. only one goroutine can enter to write. when I am writing no one can read
	defer c.rw.Unlock()
	if c.driver >= 1 {
		fmt.Println("cab is available for", name)
		time.Sleep(time.Second)
		fmt.Println("Booking confirmed", name)
		c.driver--
	} else {
		fmt.Println("cab is not available for", name)
	}
}
func main() {
	//log.New() // check logger type

	c := cab{driver: 5}

	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.getCabDrivers()
	}
	for i := 1; i <= 5; i++ {
		c.wg.Add(1)
		go c.bookCab("user" + strconv.Itoa(i))
	}


	c.wg.Wait()

}
