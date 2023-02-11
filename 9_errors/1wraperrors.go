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
		return
	}
	fmt.Println("Hello, 这是login的改动")
}

var ActionRedis = errors.New("ActionRedis: redis error")
