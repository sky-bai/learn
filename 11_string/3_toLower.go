package main

import (
	"sort"
	"strings"
)

func main() {
	strings.ToLower("ABC")
	sort.Slice(
		[]int{1, 2, 3}, func(i, j int) bool {
			return true
		})
}
