package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer func() {
			fmt.Println("defer func executed!")
			fmt.Println("recovered error == ", recover())
		}()

		for i := 0; i < 3; i++ {
			if i == 1 {
				runtime.Goexit() // 退出当前goroutine,不会报错
			}

			fmt.Println(i)
		}
	}()

	time.Sleep(2 * time.Second)
}

// 立即终止当前协程，不会影响其它协程，且终止前会调用此协程声明的defer方法。由于Goexit不是panic，所以recover捕获的error会为nil
//
//当main方法所在主协程调用Goexit时，Goexit不会return，所以主协程将继续等待子协程执行，当所有子协程执行完时，程序报错deadlock

// 计算机网路,go，lc,数据库，
