package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	key := uuid.String()
	fmt.Println(key)
	fmt.Println("1e3", 1e3)
}
