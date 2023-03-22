package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"s", "s", "r", "t", "t", "t", "q", "q", "w", "e", "c", "i", "p", "y"}

	// 统计每个元素出现的次数
	counts := make(map[string]int)
	for _, elem := range s {
		counts[elem]++
	}

	// 将出现次数排序
	type kv struct {
		Key   string
		Value int
	}
	var sortedCounts []kv
	for k, v := range counts {
		sortedCounts = append(sortedCounts, kv{k, v})
	}
	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].Value > sortedCounts[j].Value
	})

	// 打印出现次数前10的元素和次数
	fmt.Println("出现次数前10的元素和次数为：")
	for i := 0; i < 10 && i < len(sortedCounts); i++ {
		fmt.Printf("%v 出现了 %v 次\n", sortedCounts[i].Key, sortedCounts[i].Value)
	}
}
