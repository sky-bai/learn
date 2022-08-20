package main

import (
	"context"
	"fmt"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/stvp/tempredis"
)

func main() {
	server, err := tempredis.Start(tempredis.Config{})
	if err != nil {
		//panic(err)
		//fmt.Println(err)

	}
	defer server.Term()

	client := goredislib.NewClient(&goredislib.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})

	pool := goredis.NewPool(client)

	rs := redsync.New(pool)

	mutex := rs.NewMutex("test-redsync")
	ctx := context.Background()

	if err := mutex.LockContext(ctx); err != nil {
		//panic(err)
		//fmt.Println(err)
	}

	time.Sleep(time.Second * 4)
	ok, err := mutex.Extend()
	if err != nil {
		fmt.Println(err)

	}
	if !ok {
		fmt.Printf("Expected ok == true, got %v\n", ok)
	}
	time.Sleep(time.Hour)
	if _, err := mutex.UnlockContext(ctx); err != nil {

		//panic(err)
		fmt.Println(err)

	}

}
