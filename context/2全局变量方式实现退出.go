package main

import (
	"fmt"
	"sync"
	"time"
)

// if 中的break 跳出循环的判断
// continue

var exist bool
var wg2 sync.WaitGroup

func main() {
	fmt.Println(exist)
	wg2.Add(1)
	go worker1()

	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exist = true
	wg2.Wait()

	fmt.Println("over")
}

func worker1() {
	for true {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exist {
			break
		}
	}

	wg2.Done()

}
