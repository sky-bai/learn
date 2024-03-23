package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 10)
	var sendWg sync.WaitGroup
	for i := 1; i < 5; i++ {
		sendWg.Add(1)
		go func(i int) {
			defer sendWg.Done()
			//time.Sleep(time.Second * time.Duration(i))
			data := i + 10
			time.Sleep(time.Duration(data) * time.Second)
			ch <- i
		}(i)
	}
	go func() {
		sendWg.Wait()
		time.Sleep(5 * time.Second)
		close(ch)
	}()

	var wg sync.WaitGroup
	for {
		msg, ok := <-ch
		if ok {
			wg.Add(1)
			go func() {
				defer wg.Done()
				println("111", msg)
				time.Sleep(1 * time.Second)
			}()
		} else {
			fmt.Println("close")
			break
		}
	}

	wg.Wait()

	fmt.Println("over")

}
