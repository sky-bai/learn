package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	var sliceFloat = []float64{1.1, 2.2, 3.3}
	//rdb.Set(ctx, "asliceFloat", sliceFloat, 0)
	rdb.Do(ctx, "set", sliceFloat, 0)

	slice, err := rdb.Do(ctx, "get", "asliceFloat").Float64Slice()
	if err != nil {
		fmt.Println("---", err)
		return
	}
	fmt.Println("---", slice)
}
