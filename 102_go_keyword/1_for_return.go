package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker")
			doReturn()
		}
	}
}

func doReturn() {

	fmt.Println("doReturn")
	return
}

// 函数里面return不会影响上层for循环

func forSelectReturn() {
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("time.After")
			return
		}

		fmt.Println("for")
	}
}
