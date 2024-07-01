package main

import "fmt"

type Connection struct {
	waitConnection chan struct{}
}

func main() {

	// 1.读一个无缓冲只创建的channel
	readNoBufferChannel()
	// 如果没有发送者的话结果是死锁

	// 2.读一个无缓冲并关闭的channel
	readClosedNoBufferChannel()
	// 会读到channel类型的零值

}

func readNoBufferChannel() {
	c := Connection{
		waitConnection: make(chan struct{}),
	}

	data1 := <-c.waitConnection
	fmt.Println("data", data1)
}

func readClosedNoBufferChannel() {
	c := Connection{
		waitConnection: make(chan struct{}),
	}

	close(c.waitConnection)
	data1 := <-c.waitConnection
	fmt.Println("data", data1)
}
