package main

import "fmt"

func main() {
	i := 2.0
	j := 5.1
	w := j / i // 除法
	//w1 := j % i // 求余数
	fmt.Println(w)
	//fmt.Println(w1)

	u := float64(4001)
	fmt.Println(u / 1000)

	// 1.都是整数
	i1 := 2
	i2 := 5
	fmt.Println("i2 / i1 --->", i2/i1) // /求商   不是求余数
	fmt.Println("i2 % i1 --->", i2%i1) // 求余数

	// 2.都是浮点数
	j1 := 2.9
	j2 := 5.0
	fmt.Println(j2 / j1) // / 如果是浮点数 就是精度高，显示所有
	//fmt.Println(j2 % j1)

}
