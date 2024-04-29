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
	//ForRange()
	ss()
}

func ForRange() {
	v := []string{"a", "b", "c"}
	for i := range v {
		fmt.Println(i)

	}
}

func ss() {
	stri := strings.Split("22696167:113026061:1713485230073:-73:2003:1672:10700:420:20:0", ":")
	fmt.Println(len(stri))
}
