package main

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ErrRedisLockFailed = errors.New("redis lock: failed to acquire lock")
)

type Lock struct {
	rclient  redis.UniversalClient
	retries  int
	interval time.Duration
}

func New(addrs []string, db, retries int, interval time.Duration) Lock {
	if retries <= 0 {
		return Lock{}
	}
	lock := Lock{retries: retries, interval: interval}

	var password string

	parts := strings.Split(addrs[0], "@")
	if len(parts) == 2 {
		password = parts[0]
		addrs[0] = parts[1]
	}

	ropt := &redis.UniversalOptions{
		Addrs:    addrs,
		DB:       db,
		Password: password,
	}

	lock.rclient = redis.NewUniversalClient(ropt)

	return lock
}

// LockWithRetries  key redis key, unixTsToExpireNs nano time to expire
func (r Lock) LockWithRetries(key string, unixTsToExpireNs int64) error {
	for i := 0; i <= r.retries; i++ {
		err := r.Lock(key, unixTsToExpireNs)
		if err == nil {
			//成功拿到锁，返回
			return nil
		}

		time.Sleep(r.interval)
		//time.After(r.interval)
	}
	return ErrRedisLockFailed
}

func (r Lock) Lock(key string, unixTsToExpireNs int64) error {
	now := time.Now().UnixNano()
	expiration := time.Duration(unixTsToExpireNs + 1 - now)
	ctx := r.rclient.Context()

	success, err := r.rclient.SetNX(ctx, key, unixTsToExpireNs, expiration).Result()
	if err != nil {
		return err
	}

	if !success {
		v, err := r.rclient.Get(ctx, key).Result()
		if err != nil {
			return err
		}
		timeout, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		if timeout != 0 && now > int64(timeout) {
			newTimeout, err := r.rclient.GetSet(ctx, key, unixTsToExpireNs).Result()
			if err != nil {
				return err
			}

			curTimeout, err := strconv.Atoi(newTimeout)
			if err != nil {
				return err
			}

			if now > int64(curTimeout) {
				// success to acquire lock with get set
				// set the expiration of redis key
				r.rclient.Expire(ctx, key, expiration)
				return nil
			}

			return ErrRedisLockFailed
		}

		return ErrRedisLockFailed
	}

	return nil
}
