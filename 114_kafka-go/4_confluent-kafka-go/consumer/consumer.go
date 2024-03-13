package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"os/signal"
	"syscall"
)

// 需要主动全部推送 大模型的语料库

var ()

// 开6个分区

func main() {
	group := 1

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "39.108.8.92:9092",
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              group, // todo 1.group 什么时候指定
		"session.timeout.ms":    6000,
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset": "earliest", // todo 2.offset
		// Whether or not we store offsets automatically.
		"enable.auto.offset.store": false,
	})
	// 消费者组内每个消费者负责消费不同分区的数据
	// 消费者组之间互不影响

	if err != nil {
		panic(fmt.Errorf("consumer:%s,err:%v", "topic", err))
	}

	//topics := []string{"first"}
	topics := []string{"transaction"}

	err = c.SubscribeTopics(topics, nil)

	run := true

	quit := make(chan os.Signal, 1)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for run {
		select {
		case sig := <-quit:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				// Process the message received.
				fmt.Printf("Message on Topic:%s,Partition:%d,Offset:%d,Value:%s\n", *e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("Headers: %v\n", e.Headers)
				}

				// We can store the offsets of the messages manually or let
				// the library do it automatically based on the setting
				// enable.auto.offset.store. Once an offset is stored, the
				// library takes care of periodically committing it to the broker
				// if enable.auto.commit isn't set to false (the default is true).
				// By storing the offsets manually after completely processing
				// each message, we can ensure atleast once processing.
				_, err := c.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n", e.TopicPartition)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Erro11111r: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			case kafka.OffsetsCommitted:
				if e.Error != nil {
					continue
				}
			default:

				// 当数据为空字符串的时候，会走到这里

				fmt.Printf("Ignored %v\n", e)
				fmt.Printf("Type of ev: %T\n", ev)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}

func handler() {

}
