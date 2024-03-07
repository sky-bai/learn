package main

import "strings"

func main() {

	// -1 if a < b, and +1 if a > b
	i := strings.Compare("1.4.5.1", "1.4.6.0.0.1")
	println(i)
}
