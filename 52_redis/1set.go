package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var unix int64
	unix = 1670290956
	timeStr := time.Unix(unix, 0).Format("2006-01-02 15:04:05")

	fmt.Println(timeStr)
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer rdb.Close()

	val, err := rdb.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	result := rdb.HGetAll(ctx, "key")
	if result != nil {
		//sss:=result.Val()
	}

	fmt.Println(val.(string))

}
