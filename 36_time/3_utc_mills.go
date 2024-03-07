package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	utcTimeMill := time.Now().UTC().Round(time.Second).Unix()
	fmt.Println("utcTimeMill", utcTimeMill)

	utcTimeMill1 := time.Now().UTC().Unix()
	fmt.Println("utcTimeMill", utcTimeMill1)

	utcTimeMill2 := time.Now().UnixMilli()
	fmt.Println("UnixMilli", utcTimeMill2)

	str := "1970-01-01 11:11:11"

	loc, _ := time.LoadLocation("Local")

	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05 ", str, loc)

	utcTimeMil4 := the_time.UTC().Round(time.Second).UnixMilli()
	fmt.Println("utcTimeMill", utcTimeMil4)
	f1 := math.Floor(math.Floor(float64(utcTimeMil4/1000)+8*3600)/(3600*24)) + 1
	fmt.Println("f", f1)

	utcTimeMil3 := time.Now().UTC().Round(time.Second).UnixMilli()
	fmt.Println("utcTimeMill", utcTimeMil3)
	f := math.Floor(math.Floor(float64(utcTimeMil3/1000)+8*3600)/(3600*24)) + 1
	fmt.Println("f", f)

	// week
	var week int64
	week = 8 * 3600
	timeNow := time.Now().UTC().Round(time.Second).UnixMilli()
	weekSequential := math.Floor(float64(timeNow/1000+week) / 7 / 24 / 3600)
	fmt.Println("weekSequential", weekSequential)

}
