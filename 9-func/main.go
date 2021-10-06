package main

import (
	"fmt"
	"strconv"
)

func main() {
	//a, _, c := hello("", 2)
	//
	//fmt.Println(a, c)

	fmt.Println(ConvertStringtoInt("A", "100"))

}

func hello(name string, a int) (int, bool, string) {
	return 2, true, ""

}

func ConvertStringtoInt(s, x string) (string, int, error) { // err should be the last return value in your code

	a, err := strconv.Atoi(s)
	if err != nil {
		return "", 0, err
	}
	b, err := strconv.Atoi(x)
	if err != nil {
		return "", 0, err
	}

	return "sum", a + b, nil

	//return "good", a, nil
}
