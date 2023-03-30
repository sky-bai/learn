/**
Goexit
调用runtime.goExit()将立即终止当前goroutine执行，调度器
确保所有已注册defer延迟调度被执行。
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A defer go")
		func() {
			defer fmt.Println("B defer go")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")

	}() //别忘了（）
	//阻塞，防止结束
	for {
	}

	//输出
	/**
	F:\goWorkSpace\study\05协程>go run 06_runtimeGoexit.go
	B defer go
	A defer go
	*/

}
