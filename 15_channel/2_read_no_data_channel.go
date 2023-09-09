package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.创建一个有缓冲区的channel
	ch1 := make(chan int, 1024)
	ch1 <- 1
	for true {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println(data)
				time.Sleep(time.Second * 5)
			} else {
				fmt.Println("Channel Closed!")
			}
		default:
			fmt.Println("No Data!")
			time.Sleep(time.Second)
			close(ch1)
			// 如果关闭管道后 ok就是false
		}
	}
	// 取一个有缓冲的channel 会读不到数据，但是不会阻塞
}
