package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 2
	}()
	fmt.Println("time", time.Now())
	select {
	case <-ch1:
		fmt.Println("time", time.Now())
		fmt.Println("ch1")
		time.Sleep(time.Second * 5)
		fmt.Println("time", time.Now())
	case <-ch2:
		fmt.Println("time", time.Now())
		fmt.Println("ch2")

	}
	fmt.Println("time", time.Now())

	fmt.Println("end")
}
