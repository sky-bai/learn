package main

// 高阶函数 里面声明了一个变量i 低阶函数 里面对i进行自增 然后返回i

func a() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 明确高阶函数返回的是什么

func main() {
	f := a() // 返回的是低阶函数名 这个和 type test func() int 这里定义了一个函数名
	f()
	// 高阶函数返回的是低阶函数名 返回之后还需要通过 低阶函数名进行函数调用
}
