package main

import "fmt"

func main() {

	i, j := 0, 0

	for i = 0; i < 10; i++ {

		for j = 0; j < 10; j++ {
			if j == 5 {
				return
			}
		}
	}
	fmt.Println("i:", i, "j:", j)
}
