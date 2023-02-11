package main

import (
	"errors"
	"fmt"
)

func do() error {
	return ActionRedis
}

func main() {
	err := do()
	if errors.Is(err, ActionRedis) {
		fmt.Println("redis error")
	}
	fmt.Println("Hello, 这是login的改动")

	err = handle()
	if errors.Is(err, ActionRedis) { // 直接就返最底层的错误
		fmt.Println(" handle /redis error")
	}

}

var ActionRedis = errors.New("ActionRedis: redis error")

func handle() error {
	err := do()
	if err != nil {
		return err
	}
	return nil
}

// 什么时候需要打印堆栈信息？
