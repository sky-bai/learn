package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx4 = context.Background()

func main() {
	// Connect to redis.
	rdb := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer rdb.Close()
	//res, err := rdb.Get(ctx4, "sign:user1:202207").Result()
	//if err != nil {
	//}
	//fmt.Println(res)
	// 1.实现用户签到 2022 7月 offset是那天
	re, err := rdb.SetBit(ctx4, "sign:user1:202208", 1, 1).Result()
	if err != nil {
	}

	fmt.Println(re)
	//
	//res, err := rdb.GetBit(ctx4, "sign:user1:202207", 0).Result()
	//if err != nil {
	//}
	//fmt.Println(res)

	// 2.统计每月签到次数 点播次数
	count := &redis.BitCount{End: -1}
	co, err := rdb.BitCount(ctx4, "sign:user1:202207", count).Result()
	if err != nil {
	}
	fmt.Println(co)

	co1, err := rdb.StrLen(ctx4, "sign:user1:202207").Result()
	if err != nil {
	}
	// 八位等于一个字节 所以长度是 1
	fmt.Println(co1)
	ccc := "we"
	co2, err := rdb.BitOpAnd(ctx4, ccc, "sign:user1:202207", "sign:user1:202208").Result()
	if err != nil {
	}
	// 获取两个交集的个数
	fmt.Println("--2121-", co2)

	// 八位等于一个字节 所以长度是 1

}
