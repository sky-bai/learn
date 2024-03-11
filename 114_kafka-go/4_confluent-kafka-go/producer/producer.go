package producer

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
	"log"
)

// 什么时候会出现多个一个消费者组消费多个topic 动态更新缓冲区大小

// 在上层做topic的区分

// 不同topic有不同的生产者配置

const (
	INT32_MAX = 2147483647 - 1000
)

var (
	GaoDeKafkaProducer   *KafkaProducer
	TencentKafkaProducer *KafkaProducer
	TestKafkaProducer    *KafkaProducer
)

// 幂等性 保证生产的数据不会重复

type KafkaProducer struct {
	producer *kafka.Producer // 这是基础配置

	topic      string // 这是使用配置
	partition  int32
	groupId    string
	NotifyStop chan bool // 停止信号

}

func NewKafkaProducer(conf *config.KafkaConfig) *KafkaProducer {

	// 这里server读取配置文件
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.BootstrapServers, // 其实这里是集群地址 我好像不用关心消息是发完那个机器上面的 只需要关注发给那个topic就行

		//"queue.buffering.max.kbytes":   conf.QueueBuffingMaxKBytes,    // 生产者队列上允许的最大总消息大小总和 batch.size 只有数据累计到batch.size后，send才会发送给kafka,默认16kb
		//"queue.buffering.max.messages": conf.QueueBufferIngMaxMessage, // 生产者可以缓冲的最大消息数量
		// 缓冲区大小 16kb 一般不需要修改 除非分区数很多

		"batch.size":        conf.BatchSize,        // 生产者发送的消息会被分批发送，每一批的大小就是batch.size 默认16kb 可以改成32kb
		"linger.ms":         conf.LingerMs,         // 该值默认为0 来一个消息就发送 修改为5-100ms
		"compression.codec": conf.CompressionCodec, //  将数据压缩 snappy
		"acks":              conf.Acks,             // 0 1 all 0 不等待确认 1 等待leader确认 all 等待所有副本确认

		"retries": INT32_MAX,
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to write access log entry:%v", ev.TopicPartition.Error)
				}
			}
		}
	}()

	k := &KafkaProducer{
		producer:  producer,
		topic:     conf.Topic,
		partition: kafka.PartitionAny,
	}

	return k
}

type customPartitioner func(string) int32

func (k *KafkaProducer) Send(value []byte, fn customPartitioner) error {
	var msg *kafka.Message = nil

	k.partition = fn(string(value))

	// 根据规则传入的消息指定发送给对应的topic
	msg = &kafka.Message{
		// 按照key的hash code值 对 分区数 求模
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: k.partition},
		Value:          value,
	}

	err := k.producer.Produce(msg, nil)
	return err

}

func (k *KafkaProducer) Close() {
	k.producer.Flush(15 * 1000)
	k.producer.Close()
}

func (k *KafkaProducer) InitTransactions(ctx context.Context) error {
	return k.producer.InitTransactions(ctx)
}

func (k *KafkaProducer) BeginTransaction() error {
	return k.producer.BeginTransaction()
}

func (k *KafkaProducer) CommitTransaction(ctx context.Context) error {
	return k.producer.CommitTransaction(ctx)
}

func (k *KafkaProducer) AbortTransaction(ctx context.Context) error {
	return k.producer.AbortTransaction(ctx)
}

// 3.不同消息投递不同主题
// 应该是一个topic对应一个producer
// 把需要处理的方法的对象都提出来
// 队列名有什么规范吗
// 发送的数据如何清理

// 1.发送数据
// 2.指定分区发送数据
// 动态修改参数 去观察最优值

// kakfa 重启之后还是可能会有重复数据
