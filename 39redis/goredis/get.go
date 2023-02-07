package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var RedisConn *redis.Pool

func main() {
	RedisConn = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   400,
		IdleTimeout: time.Duration(1000),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				logrus.Error(err)
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	startTime := time.Now()
	Set("time111222", startTime, 0)
	fmt.Println("-----------------", time.Since(startTime))

}
func Set(key string, data interface{}, tm ...int64) error {
	conn := RedisConn.Get()
	defer conn.Close()

	switch data.(type) {
	case []byte:
		_, err := conn.Do("SET", key, data)
		if err != nil {
			return err
		}
	case string:
		_, err := conn.Do("SET", key, data)
		if err != nil {
			return err
		}
	case int:
		_, err := conn.Do("SET", key, data)
		if err != nil {
			return err
		}
	case bool:
		_, err := conn.Do("SET", key, data)
		if err != nil {
			return err
		}
	default:
		//if reflect.ValueOf(data).Kind() == reflect.Ptr { }
		value, err := json.Marshal(data)

		if err != nil {
			return err
		}

		_, err = conn.Do("SET", key, value)
		if err != nil {
			return err
		}

	}

	if tm != nil && tm[0] > 0 {
		_, err := conn.Do("EXPIRE", key, tm[0])
		if err != nil {
			return err
		}
	}

	return nil
}
