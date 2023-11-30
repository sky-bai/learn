package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

// Host n219
//
//	HostName 47.106.250.122
//	User root
func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{ // n219 n2002 n2003
		Addr:     kafka.TCP("47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"),
		Topic:    "first",
		Balancer: &kafka.LeastBytes{},
	}

	// 缓冲区大小
	w.BatchSize = 1000

	// todo 批次大小

	// linger.ms

	// 压缩

	// 1.没有指定分区 直接往topic写入数据
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Partition: 0,
			// 在 Kafka 中，Key 是消息的标识符，用于确定消息将被发送到的 partition。
			//Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Partition: 1,
			//Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Partition: 2,
			//Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("kafka failed to write messages:", err)
	}

	// 2.关闭资源
	if err := w.Close(); err != nil {
		log.Fatal("kafka failed to close writer:", err)
	}
}
