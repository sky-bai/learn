package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1[1:3])
	for index, _ := range s1 {
		s1[index] *= 2
	}
	fmt.Printf("%v\n", s1)
}
