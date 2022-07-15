package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

func main() {

	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer client.Close()

	// Create a new lock client.
	//locker := redislock.New(client)

	ctx := context.Background()
	//
	//// Try to obtain lock.
	//lock, err := locker.Obtain(ctx, "my-key", 10*time.Second, nil)
	//if err == redislock.ErrNotObtained {
	//	fmt.Println("Could not obtain lock!")
	//} else if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// Don't forget to defer Release.
	//defer lock.Release(ctx)
	//fmt.Println("I have a lock!")
	//
	//// Sleep and check the remaining TTL.
	//time.Sleep(50 * time.Second)
	//if ttl, err := lock.TTL(ctx); err != nil {
	//	//log.Fatalln(err)
	//} else if ttl > 0 {
	//	fmt.Println("Yay, I still have my lock!")
	//}
	//
	////Extend my lock.
	//if err := lock.Refresh(ctx, 10*time.Second, nil); err != nil {
	//	fmt.Println(err)
	//}
	//
	//// Sleep a little longer, then check.
	//time.Sleep(100 * time.Millisecond)
	//if ttl, err := lock.TTL(ctx); err != nil {
	//	//log.Fatalln(err)
	//} else if ttl == 0 {
	//	fmt.Println("Now, my lock has expired!")
	//}

	m := new(MyCustomLock)
	m.client = client
	lock := m.lock(ctx, "my", nil)
	time.Sleep(40 * time.Second)
	m.unLock(ctx, lock)

	time.Sleep(2 * time.Minute)

}

type MyCustomLock struct {
	client *redis.Client
}

func (m MyCustomLock) lock(ctx context.Context, key string, opt *redislock.Options) *redislock.Lock {
	locker := redislock.New(m.client)
	// Try to obtain lock.
	lock, err := locker.Obtain(ctx, key, 30*time.Second, opt)
	if err == redislock.ErrNotObtained {
		return nil
	} else if err != nil {
		return nil
	}

	// 主页
	go func() {
		myT := time.NewTicker(10 * time.Second)
		for {

			select {
			case <-ctx.Done():
				err := lock.Release(ctx)
				if err != nil {
					return
				}
				return
			case <-myT.C:
				fmt.Println("1")
				err := lock.Refresh(ctx, -1*time.Second, nil)
				if err != nil {

					return
				}
			}
		}
	}()
	return lock
}

func (m MyCustomLock) unLock(ctx context.Context, l *redislock.Lock) error {
	ctx.Done()
	err := l.Release(ctx)
	if err != nil {
		return err
	}
	return nil
}
