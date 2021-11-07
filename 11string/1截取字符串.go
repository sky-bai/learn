package main

import (
	"fmt"
	"strings"
)

func main() {
	host := "www.test.com:8080"
	host = host[0:strings.Index(host, ":")]
	fmt.Println(host)
}
