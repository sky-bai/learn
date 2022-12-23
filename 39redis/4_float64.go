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

	var divide float64 = 1000
	rdb.Set(ctx, "divide", divide, 0)
	float64Data, err := rdb.Get(ctx, "divide").Float64()
	if err != nil {
		fmt.Println("---", err)
		return
	}
	fmt.Println("-----------------", float64Data)

}
