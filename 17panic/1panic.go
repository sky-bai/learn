package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	f()
	fmt.Println("Hello, World!")
}
func f() {
	defer fmt.Println("defer f")
	panic("PANIC")
}
