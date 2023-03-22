package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	m := make(map[string]int)
	mux := sync.Mutex{}

	for i := 0; i < 1000000; i++ {
		go func() {
			randomStr := RandomStr(16)
			mux.Lock()
			if _, ok := m[randomStr]; !ok {
				m[randomStr] = 1
			} else {
				fmt.Println("重复了")
			}
			mux.Unlock()
		}()
	}
	time.Sleep(time.Second * 20)
}

// 所有样本中生成string
func RandomStr(length int) string {
	bytes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	newBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		newBytes[i] = bytes[RandomInt(0, len(bytes))]
	}

	return string(newBytes)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
