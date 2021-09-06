package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	oneMonth := time.Now().AddDate(0, -1, 0)
	startDay := now.New(oneMonth).BeginningOfMonth()
	endDay := now.New(oneMonth).EndOfMonth()
	fmt.Println("", startDay)
	fmt.Println("", endDay)
}
