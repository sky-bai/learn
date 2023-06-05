package main

import "fmt"

// 函数组合：函数组合是将多个函数组合在一起形成新的函数的过程。下面是一个示例，展示了如何使用函数组合在 Go 中实现函数的复合。
func addTwo(x int) int {
	return x + 2
}

func multiplyByThree(x int) int {
	return x * 3
}

func compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	composed := compose(addTwo, multiplyByThree)
	result := composed(4)
	fmt.Println(result) // Output: 14
}

// 在上述示例中，我们定义了两个函数 addTwo 和 multiplyByThree，它们分别将输入的整数加上 2 和乘以 3。然后，我们使用 compose 函数将这两个函数组合成一个新的函数 composed，它先应用 multiplyByThree，再应用 addTwo。最终，我们将输入值 4 传递给 composed 函数，并得到结果 14。

// 这些例子展示了如何在 Go 语言中应用函数式编程的一些基本概念，如高阶函数和函数组合。函数式编程的思想可以提高代码的模块化和可复用性，并带来更清晰和简洁的代码。虽然 Go 语言并不是纯粹的函数式编程语言，但它提供了足够的语言特性和灵活性，使我们能够以函数式风格编写代码。
