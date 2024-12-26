package main

import (
	"errors"
	"fmt"
	"io"
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

	Is()
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

func Is() {
	e1 := fmt.Errorf("e1: %w", io.EOF)
	e2 := fmt.Errorf("e2: %w + %w", e1, io.ErrClosedPipe)
	e3 := fmt.Errorf("e3: %w", e2)
	e4 := fmt.Errorf("e4: %w", e3)
	fmt.Println(errors.Is(e4, io.EOF))              // true
	fmt.Println(errors.Is(e4, io.ErrClosedPipe))    // true
	fmt.Println(errors.Is(e4, io.ErrUnexpectedEOF)) // false
}
