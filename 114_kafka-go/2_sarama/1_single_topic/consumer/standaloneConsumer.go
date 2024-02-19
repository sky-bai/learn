package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

var brokers = []string{"47.119.157.148:9092"}

func main() {
	//SinglePartition("test-z")
	Partitions("test-z")
}

// SinglePartition 单分区消费
func SinglePartition(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()
	// 参数1 指定消费哪个 topic
	// 参数2 分区 这里默认消费 0 号分区 kafka 中有分区的概念，类似于ES和MongoDB中的sharding，MySQL中的分表这种
	// 参数3 offset 从哪儿开始消费起走，正常情况下每次消费完都会将这次的offset提交到kafka，然后下次可以接着消费，
	// 这里demo就从最新的开始消费，即该 consumer 启动之前产生的消息都无法被消费
	// 如果改为 sarama.OffsetOldest 则会从最旧的消息开始消费，即每次重启 consumer 都会把该 topic 下的所有消息消费一次
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	// 会一直阻塞在这里
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}

// Partitions 多分区消费
func Partitions(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()
	// 先查询该 topic 有多少分区
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatal("Partitions err: ", err)
	}
	var wg sync.WaitGroup
	wg.Add(len(partitions))
	// 然后每个分区开一个 goroutine 来消费
	for _, partitionId := range partitions {
		go consumeByPartition(consumer, topic, partitionId, &wg)
	}
	wg.Wait()
}

func consumeByPartition(consumer sarama.Consumer, topic string, partitionId int32, wg *sync.WaitGroup) {
	defer wg.Done()
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}

func OffsetManager(topic string) {
	config := sarama.NewConfig()
	// 配置开启自动提交 offset，这样 samara 库会定时帮我们把最新的 offset 信息提交给 kafka

	// todo 什么时候该开启自动提交 offset，什么时候不该开启呢？ 对于实效性很高 但是不重要的数据 可以开启自动提交 比如 gps 点
	config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit时间间隔
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatal("NewClient err: ", err)
	}
	defer client.Close()
	// offsetManager 用于管理每个 consumerGroup的 offset
	// 根据 groupID 来区分不同的 consumer，注意: 每次提交的 offset 信息也是和 groupID 关联的
	offsetManager, err := sarama.NewOffsetManagerFromClient("myGroupID", client) // 偏移量管理器
	if err != nil {
		log.Println("NewOffsetManagerFromClient err:", err)
	}
	defer offsetManager.Close()
	// 每个分区的 offset 也是分别管理的，demo 这里使用 0 分区，因为该 topic 只有 1 个分区
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, conf.DefaultPartition) // 对应分区的偏移量管理器
	if err != nil {
		log.Println("ManagePartition err:", err)
	}
	defer partitionOffsetManager.Close()
	// defer 在程序结束后在 commit 一次，防止自动提交间隔之间的信息被丢掉
	defer offsetManager.Commit()
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Println("NewConsumerFromClient err:", err)
	}
	// 根据 kafka 中记录的上次消费的 offset 开始+1的位置接着消费
	nextOffset, _ := partitionOffsetManager.NextOffset() // 取得下一消息的偏移量作为本次消费的起点
	fmt.Println("nextOffset:", nextOffset)
	pc, err := consumer.ConsumePartition(topic, conf.DefaultPartition, nextOffset)
	if err != nil {
		log.Println("ConsumePartition err:", err)
	}
	defer pc.Close()

	for message := range pc.Messages() {
		value := string(message.Value)
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
		// 每次消费后都更新一次 offset,这里更新的只是程序内存中的值，需要 commit 之后才能提交到 kafka
		partitionOffsetManager.MarkOffset(message.Offset+1, "modified metadata") // MarkOffset 更新最后消费的 offset
	}
}
