package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("local", time.Now())
	fmt.Println("utc", time.Now().UTC())
}
