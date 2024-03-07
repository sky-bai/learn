package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// var server = "172.18.102.97:9092"
var server = "172.18.102.97:9092,172.18.102.202:9092,172.18.102.207:9092"

//var servers = []string{"172.18.102.97:9092,172.18.102.202:9092,172.18.102.207:9092"}

func main() {

	// 1.连接kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            server,
		"compression.codec":            "snappy",
		"queue.buffering.max.kbytes":   33554432, // 生产者队列上允许的最大总消息大小总和 batch.size 只有数据累计到batch.size后，send才会发送给kafka,默认16kb
		"queue.buffering.max.messages": 16384,    // 生产者可以缓冲的最大消息数量
		"linger.ms":                    1,        // 该值默认为5
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// 2.一直循环读取要发往kafka的数据
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					// 打印分区，偏移量，错误信息
				} else {
					// 处理消息
				}
			}
		}
	}()

	// 指定对应的key和value的序列化类型

	// topic一般如何创建
	// 3.异步向kafka发送消息
	topic := "first"
	Partition := kafka.PartitionAny
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: Partition},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

// 1.需要在运行的时候进行节点的上下线
// 该节点上的数据先进行迁移

// 2.如何应对lag堆积
// 5、kafka 消息丢失问题
//场景：
//
//消费端从 leader 副本 poll 了一批消息消费之后，leader 副本挂机了，之后从 ISR 选举出的副本中的消息可能是比 leader 少了的。如果此时 consumer 处理完这批数据提交 offset，消费端会丢失这部分新产生而在 kafka 中实实在在保存着的数据。
//
//解决方式：
//
//HW（high Watermark）高水位
//
//它标识了一个特定的消息偏移量（offset），消费者只能拉取到这个 offset 之前的消息。
//
//分区 ISR 集合中的每个副本都会维护自身的 LEO（Log End Offset）：俗称日志末端位移，而 ISR 集合中最小的 LEO 即为分区的 HW，对消费者而言只能消费 HW 之前的消息。
//
//附
//1.kafka 的消费组如果需要增加组员，最多增加到和 partition 数量一致，否则超过的组员只会占用资源而没有作用
//
//2.Raft 协议是啥？ 比较流行的分布式协议算法（leader 选举、日志复制）
//
//3. 分区设置：一天一亿消息大致分为 8 个分区资源可满足。
