package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	close(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1) // 读取关闭的channel 如果里面还有就取出来 没有了就取默认值

	ch2 := make(chan string, 1)
	ch2 <- "s"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2) // 读取未关闭的channel 如果里面还有就取出来 没有就会报错
	fmt.Println("---")
}
