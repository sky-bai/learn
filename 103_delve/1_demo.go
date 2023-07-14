package main

import "fmt"

// dlv 调试器

func main() {
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = i * i
	}
	fmt.Println(nums)
}
