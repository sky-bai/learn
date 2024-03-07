package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	ctx := context.Background()
	startTime := time.Now()
	fmt.Println("start", startTime)

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	var time1 = time.Now()

	rdb.Set(ctx, "time1112", time1, 0)
	time1, err := rdb.Get(ctx, "time1112").Time()
	if err != nil {
		fmt.Println("---", err)
		return
	}

	fmt.Println("-----------------", time.Since(startTime))

}
