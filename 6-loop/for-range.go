package main

import (
	"fmt"
	"os"
)

func main() {

	//for i, v := range os.Args {
	//	fmt.Println(i,v)
	//}

	//for _, v := range os.Args {
	//	fmt.Println(v)
	//}

	for i := range os.Args {
		fmt.Println(i)
	}


}
