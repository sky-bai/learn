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

	var ss sss
	ss.aa = "aa"

	rdb.Set(ctx, "struct", ss, 0)
	time, err := rdb.Get(ctx, "struct").Time()
	if err != nil {
		fmt.Println("---", err)
		return
	}

	fmt.Println("-----------------", time)
}

type sss struct {
	aa string
}
