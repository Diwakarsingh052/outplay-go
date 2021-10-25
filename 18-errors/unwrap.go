package main

import (
	"errors"
	"fmt"
)

var ErrFeesNotSubmitted error = errors.New("fees not submitted")

var ErrAdmissionCancelled error = errors.New("admission cancelled")
var ErrFoo error = errors.New("foo error")

func admission() error {
	err := fees()
	return fmt.Errorf("%w %v", err, ErrAdmissionCancelled)
}

func fees() error {
	err := foo()
	return fmt.Errorf("%w %v", err, ErrFeesNotSubmitted)
}

func foo() error {
	return fmt.Errorf("%w", ErrFoo)
}

func main() {

	err := admission()
	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
}
