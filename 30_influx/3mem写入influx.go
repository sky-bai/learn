package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"time"
)

func main() {
	p := influxdb2.NewPointWithMeasurement("memory").
		AddTag("mem", "mem").
		AddField("total", 23.2).
		AddField("available", 45.1).
		AddField("used", 45.1).
		AddField("used_percent", 45.1).
		SetTime(time.Now())
	fmt.Println("p:", p)
}
