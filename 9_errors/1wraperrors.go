package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func do() error {
	return ActionRedis
}

func main() {
	err := do()
	if errors.Is(err, ActionRedis) {
		fmt.Println("redis error")
		return
	}
	fmt.Println("Hello, 这是login的改动")
}

var ActionRedis = errors.New("ActionRedis: redis error")

func handle() error {
	err := do()
	if err != nil {
		return errors.Wrap(err, "handle")
	}
}

// 什么时候需要打印堆栈信息？
