package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// to consume messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:59092", topic, partition)
	if err != nil {
		fmt.Println("failed to dial leader:", err)
		return
	}

	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("failed to set deadline:", err)
		return
	}

	fmt.Println("start read message")

	for {
		batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

		b := make([]byte, 10e3) // 10KB max per message
		for {
			n, err := batch.Read(b)
			if err != nil {
				fmt.Println("failed to read batch:", err)
				break
			}
			fmt.Println(string(b[:n]))
			time.Sleep(1 * time.Second)
		}

		if err := batch.Close(); err != nil {
			log.Fatal("failed to close batch:", err)
		}

	}

}
