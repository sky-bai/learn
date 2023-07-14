package main

import (
	"fmt"
	"time"
)

func main() {

	input := 0
	output := 0
	closeChan := make(chan int, 1000)

	go func() {
		for {
			input++
			closeChan <- 1
		}
	}()

	go func() {
		for {
			select {
			case <-closeChan:
				output++
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Printf("input:%d,output:%d\n", input, output)
		}
	}()

	select {}

}
