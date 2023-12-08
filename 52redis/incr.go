package incr

import (
	"sync"
)

// SimpleIncrementer 是一个简单的累加器
type SimpleIncrementer struct {
	mu    sync.Mutex
	store map[string]int64
}

// NewSimpleIncrementer 创建一个新的 SimpleIncrementer 实例
func NewSimpleIncrementer() *SimpleIncrementer {
	return &SimpleIncrementer{
		store: make(map[string]int64),
	}
}

// Incr 对给定的 key 进行累加操作
func (si *SimpleIncrementer) Incr(key string, val int64) {
	si.mu.Lock()
	defer si.mu.Unlock()
	if _, ok := si.store[key]; !ok {
		si.store[key] = 0
	}

	si.store[key] += val
}

// Takeout 获取给定 key 的当前累加值并重置为0
func (si *SimpleIncrementer) Takeout(key string) int64 {
	si.mu.Lock()
	defer si.mu.Unlock()

	value := si.store[key]
	si.store[key] = 0

	return value
}
