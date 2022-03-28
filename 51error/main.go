package main

import (
	"errors"
	"fmt"
)

func te() (err error) {
	var er error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()
	raisePanic()
	return er
}
func raisePanic() {
	panic("111")
}
func main() {
	fmt.Println(te())
}
