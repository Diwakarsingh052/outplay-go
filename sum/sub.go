package sum

import "fmt"

type Store interface {
	Put(string) error
}

func sub() {
	fmt.Println("I am calling sub")

}
