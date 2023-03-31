package main

import "fmt"

func main() {
	camera := 4
	// 1.获取文件数量
	fileTotal := 0
	for camera > 0 {
		if camera&1 == 1 {
			fileTotal++
		}
		camera = camera >> 1
		fmt.Println("camera", camera)
	}
	fmt.Println("------------")

	yu()
	huo()
	fei()
	zuoYi()
	youYi()
}

func yu() {
	// 在 Go 中，& 操作符用来在两个整数之间进行位 AND 运算。如果两个相应位都是 1，则该位的结果为 1，否则为 0。
	// 可以通过 & 来判断一个数字是奇数还是偶数,我们可以将数字和值 1 使用 & 做 AND 运算。如果结果是 1，那说明原来的数字是一个奇数。	// 如果结果是 0，那说明原来的数字是一个偶数。
	i := 4
	fmt.Println("位与&&&&")
	fmt.Println(i & 1)
}

func huo() {
	i := 1
	fmt.Println("位或||||")
	fmt.Println(i | 1)
}

func fei() {
	i := 1
	fmt.Println("异或^^^^")
	fmt.Println(^i)
}
func zuoYi() {
	i := 1
	i = i << 1
	fmt.Println("左移<<")
	fmt.Println(i)
}
func youYi() {
	i := 8
	i = i >> 1
	fmt.Println("右移>>")
	fmt.Println(i)
	// prints:
	// 01111000
	// 00111100
	// 00011110
}

// &   位与
// |   位或
// ^   异或
// &^   位与非
// <<   左移
// >>   右移
