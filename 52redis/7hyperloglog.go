package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to redis.
	rdb := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer rdb.Close()

	ctx7 := context.Background()
	for i := 0; i < 10; i++ {
		if err := rdb.PFAdd(ctx7, "myset", fmt.Sprint(i)).Err(); err != nil {
		}
	}

	card, err := rdb.PFCount(ctx7, "myset").Result()
	if err != nil {
	}

	fmt.Println("set cardinality", card)

}
