package main

import "fmt"

// 高阶函数处理一个函数 这个函数的入参是字符串的s对象
// 低阶函数处理一个字符串的s对象

// 低阶函数调用高阶函数里面的函数处理低阶函数

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("started")
		f(s)
		fmt.Println("done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	f := decorator(Hello)
	f("hello,world")
}

// 从使用者的角度来看 我需要传入一个函数名 和 入参
// 之前是过程式编程 直接调用函数 传入函数的入参

// 代办事项 查看net包的test
// 前端部署
