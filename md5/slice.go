package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
		}
		a[k] = 100 + v
	}
	fmt.Println(a)
}

// 数组是拷贝的值是赋给 下面的值 切片是引用 000000
