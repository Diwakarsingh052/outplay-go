package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	var key string = "anyKey" // bad practice
	ctx = context.WithValue(ctx, key, "1000")
	fetchValue(ctx, key)
}

func fetchValue(ctx context.Context, k string) {

	v := ctx.Value(k)

	//if v != nil {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("not there")
	//}

	a, ok := v.(int64)
	if !ok {
		fmt.Println("not there")
		return
	}
	fmt.Println(a)

}
