package main

import (
	"fmt"
	"time"
)

type esClient struct {
	logDataChan chan string
}

var es1 esClient

func main() {
	es1 = esClient{
		logDataChan: make(chan string, 10),
	}
	es1.logDataChan <- "hellosdff"
	for i := 0; i < 10; i++ {
		go sendToES(i)
	}
	es1.logDataChan <- "hellosdff"
	es1.logDataChan <- "hellosdff"
	es1.logDataChan <- "hellosdff"

	time.Sleep(time.Second * 10)
}

func sendToES(i int) {
	for msg := range es1.logDataChan {
		fmt.Println(i, msg)
	}
}
