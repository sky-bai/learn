package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Get(60 * 60 * 24))
	fmt.Println(fmtDuration(60 * 60 * 24))
}

func Get(i int) string {
	h := i / 3600
	m := i % 3600 / 60
	s := i % 3600 % 60
	fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	p := ""
	//if h > 0 {
	p += fmt.Sprintf("%dh", h)
	//}
	//if m > 0 {
	p += fmt.Sprintf("%dm", m)
	//}
	//if s > 0 {
	p += fmt.Sprintf("%ds", s)
	//}
	return p
}

func fmtDuration(i int) string {
	duration := time.Duration(i) * time.Second
	return duration.String()
}
