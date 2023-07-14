package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("start:", t)
OuterLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
		for i := 0; i < 3; i++ {
			fmt.Println("break", time.Now().Sub(t))
			break OuterLoop
		}
	}
	fmt.Println("end")
}

//
