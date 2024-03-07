package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("退出")
				return

			}
		}

	}()
	close(c)
	time.Sleep(time.Second * 20)
}
