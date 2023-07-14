package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func main() {
	ForRange()
}

func ForRange() {
	v := []string{"a", "b", "c"}
	for i := range v {
		fmt.Println(i)

	}
}
