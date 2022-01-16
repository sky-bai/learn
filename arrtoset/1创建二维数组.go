package main

import "fmt"

func main() {
	x := 2
	y := 4

	table := make([][]int, x)
	for i, v := range table {
		fmt.Println("i:", i)
		fmt.Println("v:", v)
		table[i] = make([]int, y)
	}

	fmt.Println("The size of the table: ", len(table), "x", len(table[0]))
	fmt.Println("The value of table: ", table)

	// 创建一个行数为4的二维数组
	arr := make([][]int, 4)
	// 创建一个列数为3的二维数组
	for i := range arr {
		arr[i] = make([]int, 4)
	}
	fmt.Println("The size of the arr: ", len(arr), "x", len(arr[0]))
	fmt.Println("The value of arr: ", arr)
}
