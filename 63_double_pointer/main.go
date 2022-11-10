package main

import "fmt"

func main() {
	var m1 = []int{1, 2, 3}
	var m2 = []int{0, 0, 1, 1, 2, 2, 3, 5, 8}
	var m3 = []int{}
	flag := 0
	up, down := 0, 0
	for up < len(m1) {
		if m1[up] == m2[down] {
			flag = 1
			m3 = append(m3, m1[up])
			down++
			continue
		} else {
			if flag == 1 {
				flag = 0
				up++
				continue
			} else {
				down++
				continue
			}
		}
	}
	fmt.Println(m3)

}
