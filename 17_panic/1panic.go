package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("父协程捕获子协程的err:", err)
		}
	}()
	f()
}
func f() {

	panic("子协程 PANIC")
}

// 这个函数本身要panic,就在这个函数内recover  这样就会使main函数继续执行
