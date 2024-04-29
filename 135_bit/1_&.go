package main

import "fmt"

func main() {
	if 1&1 == 1 {
		fmt.Println("1 & 1 == 1")
	}

	if 2&1 == 1 {
		fmt.Println("2 & 1 == 1")
	}

	if 3&1 == 1 {
		fmt.Println("3 & 1 == 1")
	}
}
