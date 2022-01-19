package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
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
	//
	//// write line protocol
	//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	//// Flush writes
	//writeAPI.Flush()

	// get non-blocking write client
	//writeAPI := client.WriteAPI("test", "test")

	//p := influxdb2.NewPoint("stat",
	//	map[string]string{"unit": "temperature"},
	//	map[string]interface{}{"avg": 24.5, "max": 45},
	//	time.Now())
	//// write point asynchronously
	//writeAPI.WritePoint(p)
	//create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.1).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
	//Get query client
	queryAPI := client.QueryAPI("test")

	query := `from(bucket:"test")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		fmt.Println("field", result.Record().Field())
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}

}
