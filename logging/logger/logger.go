package logger

import (
	"fmt"
	"io"
)

type logger struct {
	ch chan string
}

func New(w io.Writer, cap int) *logger {

	l := logger{
		ch: make(chan string, cap),
	}

	go func() {
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}

	}()
	return &l
}

func (l *logger) Println(v string) {

	select {
	case l.ch <- v:
	default:
		fmt.Println("drop")
	}

}
