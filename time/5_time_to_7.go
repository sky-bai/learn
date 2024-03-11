package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 7, 0, 0, 0, now.Location())
	// 计算时间差
	diff := now.Sub(today)
	fmt.Println("---", diff)
}
