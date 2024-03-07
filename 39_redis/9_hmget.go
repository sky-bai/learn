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
		DB:   0,
	})

	// 关闭客户端连接
	defer rdb.Close()

	rdb.HMSet(ctx, "hash", "key1", "value1", "key2", "value2", "key3", "value3")

	data, err := rdb.HMGet(ctx, "hash", "key1", "key4", "key3").Result()
	if err != nil {
		fmt.Printf("HMGet err:%s\n", err)
	}

	fmt.Println(data)

	for _, v := range data {
		if xx, ok := data.(redis.Nil); ok {
			fmt.Println("nil111")
		} else {
			fmt.Println(v)
		}
	}
}
