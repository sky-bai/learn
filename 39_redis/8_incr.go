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
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	// 关闭客户端连接
	defer rdb.Close()

	//// 1.incr
	//data, err := rdb.Incr(ctx, "incr").Result()
	//if err != nil {
	//	panic(err)
	//}
	//println(data)
	//
	//// 2.incrBy
	//data1, err := rdb.IncrBy(ctx, "incr", 3).Result()
	//if err != nil {
	//	panic(err)
	//}
	//println(data1)

	// 3.两个方法结果都是累加后的值

	// 4.对一个key进行incrBy操作，将获取到累加后的值进行判断，如果大于20，就将值重置为0
	// Lua中的tonumber函数可以将一个字符串或数字类型的值转换为一个数值类型。
	// 定义Lua脚本
	script := `
    local key = KEYS[1]
	local increment = tonumber(ARGV[1])
	local currentVal = redis.call('INCRBY', key, increment)
	
	if currentVal > 20 then
		redis.call('SET', key, 0)
	end
	
	return currentVal
    `

	// key名称
	key := ""

	// 执行Lua脚本
	result, err := rdb.Eval(ctx, script, []string{key}, 10).Int64()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func YiKa() {

}
