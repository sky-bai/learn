package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`foo.*`, []byte(`seafod`))
	fmt.Println(matched, err)
}
