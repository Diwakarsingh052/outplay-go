package main

import (
	"fmt"
	"sort"
)

type student struct {
	name  string
	marks []int
}

func main() {
	var keys []int

	class := map[int]student{
		1: {
			name:  "abc",
			marks: []int{10, 60, 60},
		},
		2: {
			name:  "xyz",
			marks: []int{10, 60, 60},
		},
		3: {
			name:  "efgh",
			marks: []int{10, 60, 60},
		},
	}
	keys = append(keys, 2, 1, 3)
	sort.Ints(keys)

	//for k, v := range class {
	//	fmt.Println(k, v)
	//	fmt.Println(class[k])
	//}

	for _, k := range keys {

		fmt.Println(k, class[k])
	}

}
