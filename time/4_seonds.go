package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1.Minute(), t1.Second())
}
