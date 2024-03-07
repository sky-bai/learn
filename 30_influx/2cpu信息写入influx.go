package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func main() {

	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://121.196.163.8:8086", "BLt2XFaQv7b1dI5X8syeNQJPpISJb_mJly4DAxAVLTy7q_FvWnXaDu7cgfh6GkiTMKggjWiPJiu53NHXP4vGBQ==")
	// always close client at the end
	client.QueryAPI("test")
	defer client.Close()

	//// get non-blocking write client
	writeAPI := client.WriteAPI("test", "test")

	ticker := time.Tick(time.Second)
	for {
		select {
		case <-ticker:
			percent, _ := cpu.Percent(time.Second, false)
			fmt.Println("cpu使用率:", percent[0])

			p := influxdb2.NewPointWithMeasurement("cpuInfo").
				AddTag("unit", "cpu").
				AddField("percent", percent[0]).
				SetTime(time.Now())
			// write point asynchronously
			writeAPI.WritePoint(p)
			// Flush writes
			writeAPI.Flush()
		}
	}

}
