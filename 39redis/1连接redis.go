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
	reply, err := con.Do("SET", "value", "ss", "EX", "100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
