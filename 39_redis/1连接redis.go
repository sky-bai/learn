package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	var pool = &redis.Pool{
		MaxActive:   2000,
		IdleTimeout: 20,
		MaxIdle:     10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	con := pool.Get()
	reply, err := con.Do("SET", "value", 1, "EX", "100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
	do, err := con.Do("GET", "value")
	if err != nil {
		fmt.Println("error111")
		return
	}
	if value, ok := do.(int); ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}
}

// 针对不同类型的数据都可以设置进去，取出来也是还原到设置之前的类型
