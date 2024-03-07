package main

import "fmt"

func main() {
	// 该类型放多少个
	messageChannel := make(chan string, 2)
	fmt.Println("channel的长度", len(messageChannel))
	fmt.Println("channel的容量", cap(messageChannel))
	messageChannel <- "hello"
	fmt.Println("channel的长度", len(messageChannel))
	fmt.Println("channel的容量", cap(messageChannel))
	messageChannel <- "helloaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	fmt.Println("channel的长度", len(messageChannel))
	fmt.Println("channel的容量", cap(messageChannel))

}
