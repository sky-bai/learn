package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := sort.Search(len(a), func(i int) bool { return a[i] >= 30 })
	fmt.Println(b)
}
