package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	urlStr := "http://res.spreadwin.cn/803278214323710m9lsUPNz9Q8RCX8/111.zip"
	parse, err := url.Parse(urlStr)
	if err != nil {
		fmt.Print(err)
		return
	}
	str := strings.Split(parse.Path, ".")
	fmt.Println(parse.Path)
	fmt.Println(str[0])
}
