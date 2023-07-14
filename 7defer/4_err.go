package main

import (
	"fmt"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)
		}
	}()
	//panic("panicli")
	//pan()
	//
	//time.Sleep(10 * time.Second)

	fmt.Println("---", 10&253)
	// 10  1010
	//     1101

}

func pan() {
	panic("panic")
}
