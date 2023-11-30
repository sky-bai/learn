package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer rdb.Close()

	result, err := rdb.SMembers(ctx, "key").Result()
	if err != nil {
		log.Println("1111", err)
		return
	}
	fmt.Println(result)
}
