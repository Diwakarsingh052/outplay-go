package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Func string
	Err  error
}

func (q *QueryError) Error() string {
	return q.Func + ":" + q.Err.Error()
}

func main() {
	//_, err := strconv.Atoi("a")
	//os.OpenFile()
	var e *QueryError // nil

	err := SearchSomething()

	if errors.As(err, &e) {
		fmt.Println("true ")
		fmt.Println(e.Func)
	} else {
		fmt.Println("false")
		fmt.Println(e)
	}

	//fmt.Println(err)

	//var err error = errors.New("any error")
	//fmt.Printf("%T", err)

}

func SearchSomething() error {
	// do your stuff
	return &QueryError{
		Func: "SearchSomething",
		Err:  errors.New("not able to find what you are searching"),
	}
}
