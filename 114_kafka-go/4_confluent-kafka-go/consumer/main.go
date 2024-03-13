package consumer

import (
	"fmt"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	TestKafkaConsumer *KafkaConsumer
)

// 一个分区 三个副本
// 什么时候会从一个分区扩大到多个分区 分区数只能增加不能减少 为什么不能减少昵
// 需要指定消费位置
// 一台机器有topic的多个分区 不同机器存储分区的副本
func main() {

	consumer.Consumer(HandlerTest)

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	consumer.Close()

}

type KafkaConsumer struct {
	consumer *kafka.Consumer
	groupId  string
}

func NewKafkaConsumer(conf *config.KafkaConsumerConfig) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          conf.GroupId,
		"auto.offset.reset": autoOffsetReset,
	})

	if err != nil {
		panic(fmt.Errorf("consumer:%s,group:%d,err:%v", "topic", conf.GroupId, err))
	}

	err = c.SubscribeTopics(conf.Topic, nil)
	if err != nil {
		panic(fmt.Errorf("consumer:%s,group:%d,err:%v", "topic", conf.GroupId, err))
	}

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

}

func (k *KafkaConsumer) Consumer(fn func([]byte)) {
	// 3.读取
	for {
		msg, err := k.consumer.ReadMessage(-1)
		if err == nil {
			// 处理结果
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			fn(msg.Value)
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (k *KafkaConsumer) Close() {
	k.consumer.Close()
}

func HandlerTest(msg []byte) {
	fmt.Println("handler msg:", string(msg))
}

func Demo() {
	// 1.连接
	//c, err := kafka.NewConsumer(&kafka.ConfigMap{
	//	"bootstrap.servers": servers,
	//	"group.id":          myGroup,
	//	"auto.offset.reset": autoOffsetReset,
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 2.订阅
	//c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	//
	//// A signal handler or similar could be used to set this to false to break the loop.
	//run := true
	//
	//// 3.读取
	//for run {
	//	msg, err := c.ReadMessage(time.Second)
	//	if err == nil {
	//		// 处理结果
	//		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	//	} else if !err.(kafka.Error).IsTimeout() {
	//		// The client will automatically try to recover from all errors.
	//		// Timeout is not considered an error because it is raised by
	//		// ReadMessage in absence of messages.
	//		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	//	}
	//}
	//
	//c.Close()
}
