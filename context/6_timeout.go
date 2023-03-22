package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("time", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded" }
	}
	fmt.Println("time", time.Now())
}
