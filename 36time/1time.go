package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02"))
	currentTime := time.Now()
	yesterTime := currentTime.Add(-24 * time.Hour)
	fmt.Println(yesterTime)
}
