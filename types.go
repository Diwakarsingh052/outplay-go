package main

import (
	"fmt"
)

type money int64 // money is a new type // it's not an alias // and it is a type like float, bool

//type id int64

//var id int64

func main() {

	var rupee money = 100
	//var a, b, c = "", true, 10

	var a int64 = int64(rupee)
	//time.Duration
	fmt.Println(a)

}
