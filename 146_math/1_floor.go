package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	rate := 33
	amount := math.Floor(float64(x * rate))

	fmt.Println(amount)
}
