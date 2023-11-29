package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {

	config := sarama.NewConfig()

	// 缓冲区大小
	// batch.size 只有数据累计到batch.size后，send才会发送给kafka,默认16kb
	config.Producer.Flush.Bytes = 33554432
	// 批次大小
	config.Producer.Flush.Messages = 16384
	// linger.ms
	// 默认0，表示数据必须立即发送，>0表示数据在linger.ms后send才发送
	config.Producer.Flush.Frequency = 1
	// 压缩
	config.Producer.Compression = sarama.CompressionSnappy
	// 设置ack 为1 leader和follower都确认
	config.Producer.RequiredAcks = sarama.WaitForLocal
	// 重试次数
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewAsyncProducer([]string{"47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"}, config)
	if err != nil {
		fmt.Println("kafka Failed to start consumer: ", err)
		return
	}

	producer.Input() <- &sarama.ProducerMessage{
		Topic: "first",
		Value: sarama.StringEncoder("test"),
	}

	err = producer.Close()
	if err != nil {
		fmt.Println("kafka Failed to close producer: ", err)
		return
	}

}

// 0:生产者发送过来的数据，不需要等数据罗盘应答
// 1:生产者发送过来的数据，需要等待leader应答，不需要等待follower应答
// -1:生产者发送过来的数据，需要等待leader应答，需要等待follower应答
