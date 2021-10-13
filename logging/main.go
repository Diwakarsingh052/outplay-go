package main

import (
	"fmt"
	"logging/logger"
	"os"
	"os/signal"
	"time"
)

type disk struct {
	diskFull bool
}

func (d *disk) Write(data []byte) (n int, err error) {

	for d.diskFull { // ==true
		time.Sleep(time.Second)
	}

	fmt.Println(string(data))
	return len(data), nil

}

func main() {
	var d disk
	//l := log.New(&d, "", log.LstdFlags)
	l := logger.New(&d, 10)
	for i := 0; i < 10; i++ {
		go func(id int) {

			for {

				l.Println(fmt.Sprintf("%d: log data", id))

			}

		}(i)

	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.diskFull = !d.diskFull

	}

}
