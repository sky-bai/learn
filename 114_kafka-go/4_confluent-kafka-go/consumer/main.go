package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var servers = ""
var myGroup = ""

func main() {

	// 1.连接
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          myGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	// 2.订阅
	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	// 3.读取
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
