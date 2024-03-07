package main

import (
	"fmt"
	sync2 "sync"
)

// defer 是在函数退出的时候进行执行的
// 函数 return 之前处理某个语句或函数

func main() {
	sync := sync2.Mutex{}
	for i := 0; i < 3; i++ {
		sync.Lock()
		defer sync.Unlock()
		fmt.Println(1)
	}
}

func main1() {
	sync := sync2.Mutex{}
	for i := 0; i < 3; i++ {
		go func() {
			sync.Lock()
			defer sync.Unlock()
			fmt.Println(2)
		}()
	}
}
