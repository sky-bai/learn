package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	var time = time.Now()

	rdb.Set(ctx, "time", time, 0)
	time, err := rdb.Get(ctx, "time").Time()
	if err != nil {
		fmt.Println("---", err)
		return
	}

	fmt.Println("-----------------", time)

}
