package main

import "fmt"

func main() {

	f()
	fmt.Println("Hello, World!")
}
func f() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	defer fmt.Println("defer f")
	panic("PANIC")
}

// 这个函数本身要panic,就在这个函数内recover  这样就会使main函数继续执行
