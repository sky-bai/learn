package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/spf13/cast"
	"time"
)

func main() {
	time.Now()

	str := "2022-03-11"

	loc, _ := time.LoadLocation("Local")

	the_time, _ := time.ParseInLocation("2006-01-02", str, loc)
	fmt.Println("11111", now.With(the_time).BeginningOfMonth())

	xxxTime := int(now.With(the_time).Month())
	fmt.Println()
	if xxxTime < 9 {
		x := cast.ToString(xxxTime)
		x = "0" + x
		xxxTime = cast.ToInt(x)
	}

	fmt.Println("000", now.With(the_time).Day())
	s, _ := now.With(the_time).Parse("1999-12-12")
	fmt.Println("---", cast.ToInt(s))

	fmt.Println(int(the_time.Month()))
	fmt.Println(int(the_time.Month()))
	fmt.Println(now.With(the_time).BeginningOfYear())

}
