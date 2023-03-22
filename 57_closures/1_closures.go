package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values { // for 迭代器 已经修改了v的值,因为是遍历变量v，遍历完了，v的值就是最后一个元素的值
		go func(u string) { // 这里是行参
			fmt.Println(u) // 闭包共享v这个变量
			done <- true
		}(v) // 这里是实参
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}

// 并发里面使用闭包的时候，需要注意的是，闭包里面的变量是共享的，所以在并发的时候，需要注意闭包里面的变量的值是否会被修改，如果会被修改，那么需要注意并发安全的问题。

// 是的，当 goroutine 执行时，循环变量 v 的值已经被修改成了最后一个元素 "c"。这是因为 goroutine 是异步执行的，for 循环会继续执行并递增 v，直到它的最终值为 "c"。在此之后，
// 所有 goroutine 都在闭包中引用了 v 的最终值 "c"。因此，当它们被执行时，都会打印出 "c"，而不是它们分配时期望的值。
// 为了避免这种情况，我们可以将循环变量作为 goroutine 的参数进行传递，这样每个 goroutine 将使用它自己的变量值而不会与其他 goroutine 共享相同的变量。
