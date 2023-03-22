package main

import (
	"fmt"
)

func main() {
	total := 101
	size := 10
	times := total/size + 1
	ch := make(chan int, times)
	for i := 0; i < times; i++ {
		ch <- i + 1
		go func() {
			page := <-ch
			fmt.Println("当前page: ", page)
		}()
	}
	// 优雅的退出

}
