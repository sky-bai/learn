package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var servers = ""
var myGroup = ""
var autoOffsetReset = ""

// 一个分区 三个副本
// 什么时候会从一个分区扩大到多个分区 分区数只能增加不能减少 为什么不能减少昵
// 需要指定消费位置
// 一台机器有topic的多个分区 不同机器存储分区的副本
func main() {

	// 1.连接
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          myGroup,
		"auto.offset.reset": autoOffsetReset,
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
			// 处理结果
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
