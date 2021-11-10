package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 让leader 和 follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	msg := &sarama.ProducerMessage{}
	msg.Topic = "白"
	msg.Value = sarama.StringEncoder("白")

	producer, err := sarama.NewSyncProducer([]string{"0.0.0.0:9092"}, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer producer.Close()
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("pid : %v,offset: %v\n\n", pid, offset)
	fmt.Println("执行完成")

}
