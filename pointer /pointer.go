package main

import "fmt"

func main() {
	// *int 指针变量 就是指向一个变量的地址 p 就是地址 *类型就是指针 就是地址 *指针变量就是找到改指针指向的变量
	var p *int   // var 只是声明该变量 并没有分配内存地址
	p = new(int) // new 是返回类型变量的地址
	*p = 1
	fmt.Println(p, &p, *p)
}
