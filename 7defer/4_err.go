package main

import (
	"fmt"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)
		}
	}()
	panic("panicli")
	pan()

	time.Sleep(10 * time.Second)

}

func pan() {
	panic("panic")
}
