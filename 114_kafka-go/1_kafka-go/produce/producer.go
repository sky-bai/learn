package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	topic := "first"
	partition := 0

	// n2002 47.119.157.148
	// 1.连接集群
	conn, err := kafka.DialLeader(context.Background(), "tcp", "47.119.157.148:9092", topic, partition)
	if err != nil {
		fmt.Println("kafka failed to dial leader:", err)
		return
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
