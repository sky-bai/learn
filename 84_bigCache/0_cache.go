package main

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
	DelOldest()
	Len() int
}

// 考虑并发（即单goroutine读写）和GC问题

// fifo 先进先出
// lru 最近最少使用

// 本地内存和redis缓存的同步
