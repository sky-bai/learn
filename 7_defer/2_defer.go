package main

import "fmt"

func main() {
	normalReturn := false
	panicFunc := func() {
		panic("panic")
	}

	defer func() {
		println("defer 3")
		fmt.Println("normalReturn", normalReturn)
	}()

	func() {
		defer func() {
			println("inside defer 2")
			if !normalReturn {
				if r := recover(); r != nil {
					println("inside recover 2")
					fmt.Println("recover", r) // r 就是错误信息
				}
			}
		}()

		println("main 1")
		panicFunc()
		normalReturn = true
	}()

	fmt.Println("normalReturn", normalReturn)
}
