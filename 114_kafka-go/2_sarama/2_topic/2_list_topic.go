package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
		return
	}

	// 1.拿到指定分区下的分区列表
	partitionList, err := consumer.Partitions("first") // 返回每个分区的标识
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	//var wg sync.WaitGroup
	// 为每一个分区创建一个对应的分区消费者
	for partitionID := range partitionList {
		fmt.Println("partitionID: ", partitionID)
		//wg.Add(1)
		//pc, err := consumer.ConsumePartition("web_log", int32(partitionID), 2_sarama.OffsetNewest)
		//if err != nil {
		//	fmt.Println("Failed to start consumer for partition ", partitionID, ": ", err)
		//	return
		//}
		//defer pc.AsyncClose()
		//go func(2_sarama.PartitionConsumer) {
		//	for msg := range pc.Messages() {
		//		fmt.Printf("Partition:%d,offset:%d key:%s value:%s ", msg.Partition, msg.Offset, msg.Key, msg.Value)
		//	}
		//}(pc)
	}
	//wg.Wait()
	// 创建client
	newClient, err := sarama.NewClient([]string{"47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 获取所有的topic
	topics, err := newClient.Topics()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("topics", topics)

	// __consumer_offsets 是 Kafka 中内置的系统主题之一，它用于存储消费者组的偏移量信息。这个主题是 Kafka 用来维护消费者组的消费进度的，记录了每个消费者组在每个分区上消费的偏移量。这些偏移量是用来跟踪每个消费者组在每个主题的每个分区上的消费进度，确保每个消费者能够从上一次中断的位置继续消费消息。
	// 该主题的数据结构是特殊的，并且由 Kafka 内部管理和维护，存储在 Kafka 的内部日志（topic log）中，一般情况下用户无需直接操作该主题。
}
