package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	PrintData(ctx, 2*time.Second, "hello how are you")

}

func PrintData(ctx context.Context, d time.Duration, input string) { // ctx should be your first param

	select {
	case <-time.After(d):
		fmt.Println(input)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	//abc(ctx)
}

//func abc(ctx context.Context) {
//
//}
