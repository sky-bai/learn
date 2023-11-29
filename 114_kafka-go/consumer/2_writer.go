package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{ // n219 n2002 n2003
		Addr:     kafka.TCP("47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"),
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{},
	}
	// 如果指定分区 这里的key是什么
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			// 在 Kafka 中，Key 是消息的标识符，用于确定消息将被发送到的 partition。
			//Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			//Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			//Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
