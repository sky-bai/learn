package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`foo.*`, []byte(`seafod`))
	fmt.Println(matched, err)
	matched1, err := regexp.Match(`^[a-zA-Z0-9_]{6,128}$`, []byte("4512w34"))
	fmt.Println(matched1, err)
}
