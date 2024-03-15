package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
)

var (
	TestKafkaConsumer *KafkaConsumer
)
var (
	ConsumerDoneChannel = make(chan struct{})
)

// 一个分区 三个副本
// 什么时候会从一个分区扩大到多个分区 分区数只能增加不能减少 为什么不能减少昵
// 需要指定消费位置
// 一台机器有topic的多个分区 不同机器存储分区的副本

type KafkaConsumer struct {
	consumer *kafka.Consumer
	topic    []string
	groupId  string
	doneChan chan struct{}
}

func NewKafkaConsumer(conf *config.KafkaConsumerConfig) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.BootstrapServers,
		"group.id":          conf.GroupId,
		"auto.offset.reset": conf.AutoOffsetReset,
	})

	if err != nil {
		panic(fmt.Errorf("consumer:%s,group:%d,err:%v", "topic", conf.GroupId, err))
	}

	err = c.SubscribeTopics(conf.Topic, nil)
	if err != nil {
		panic(fmt.Errorf("consumer:%s,group:%d,err:%v", "topic", conf.GroupId, err))
	}

	//// A signal handler or similar could be used to set this to false to break the loop.
	//run := true

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

	return &KafkaConsumer{
		consumer: c,
		groupId:  conf.GroupId,
		topic:    conf.Topic,
		doneChan: make(chan struct{}),
	}

}

func (k *KafkaConsumer) Consumer(fn func(*kafka.Message)) {
	// 3.读取
	for {
		select {
		case <-k.doneChan:
			fmt.Println("消费者停止消费")
			ConsumerDoneChannel <- struct{}{}
			return
		default:
			fmt.Println("消费者开始消费")
			msg, err := k.consumer.ReadMessage(-1)
			if err == nil {
				// 处理结果
				//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				fn(msg)
			} else if !err.(kafka.Error).IsTimeout() {
				// The client will automatically try to recover from all errors.
				// Timeout is not considered an error because it is raised by
				// ReadMessage in absence of messages.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}

	}
}

func (k *KafkaConsumer) Done() {
	fmt.Printf("333")
	k.doneChan <- struct{}{}
}

func (k *KafkaConsumer) Close() {
	k.consumer.Close()
}

func HandlerTest(msg *kafka.Message) {
	fmt.Printf("Consumer Message on Topic:%s,Partition:%d,Offset:%d,Value:%s\n", *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset, string(msg.Value))
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
