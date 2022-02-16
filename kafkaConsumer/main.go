package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"121.196.163.8:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
		return
	}

	// 拿到指定分区下的分区列表
	partitionList, err := consumer.Partitions("web_log") // 返回每个分区的标识
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	var wg sync.WaitGroup
	// 为每一个分区创建一个对应的分区消费者
	for partitionID := range partitionList {
		fmt.Println("partitionID: ", partitionID)
		wg.Add(1)
		pc, err := consumer.ConsumePartition("web_log", int32(partitionID), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition ", partitionID, ": ", err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d,offset:%d key:%s value:%s ", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
