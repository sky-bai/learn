package main

func main() {
	// 匿名函数
	// 匿名函数是没有名字的函数，匿名函数的类型是函数类型，可以赋值给变量，也可以作为参数传递给其他函数。
	// 匿名函数可以直接调用，也可以作为函数值赋值给变量，最后通过变量来调用。
	// 匿名函数可以在函数内部定义，也可以在函数外部定义。
	// 匿名函数的语法格式如下：
	// func(参数)(返回值){
	// 	函数体
	// }(实参)

	defer func() {
		println("hello world defer")
	}()
	println("hello world")
	func() {
		println("hello world func")
	}()

}

// go语言函数式编程的使用场景
