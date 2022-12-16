package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var divide int = 1000
	var divi int = 1
	fmt.Println("--------", divide/divi)

	var pipucot int = 1
	flag.IntVar(&pipucot, "pc", 1000, "pip批量处理个数")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	ctx := context.Background()

	pip := rdb.Pipeline()

	t := time.Now()
	cmdcot := 100000

	for k := 1; k < cmdcot; k++ {

		key := fmt.Sprint("keypip%d", k)
		err := pip.Set(ctx, key, k, 0).Err()
		if err != nil {
			panic(err)
		}
		if k%pipucot == 0 {
			fmt.Println("k:", k)
			fmt.Println("pipucot:", pipucot)
			pip.Exec(ctx)
		}
	}
	println("pip use time:", time.Since(t).Milliseconds())

	t1 := time.Now()
	for k := 1; k < cmdcot; k++ {

		key := fmt.Sprint("key%d", k)
		err := rdb.Set(ctx, key, k, 0).Err()

		if err != nil {
			panic(err)
		}
	}
	// 清除所有的key
	rdb.FlushDB(ctx)
	println("no pip use time:", time.Since(t1).Milliseconds())

}
