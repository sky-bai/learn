package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
	"log"
	"time"
)

// 什么时候会出现多个一个消费者组消费多个topic 动态更新缓冲区大小

// 在上层做topic的区分

// 不同topic有不同的生产者配置

var (
	GaoDeKafkaProducer   *KafkaProducer
	TencentKafkaProducer *KafkaProducer
	TestKafkaProducer    *KafkaProducer
)

func init() {

	GaoDeKafkaProducer = NewKafkaProducer(&config.KafkaConfig{
		LingerMs:                 config.GaoDeConfigSetting.LingerMs,
		BatchSize:                config.GaoDeConfigSetting.BatchSize,
		QueueBuffingMaxKBytes:    config.GaoDeConfigSetting.QueueBuffingMaxKBytes,
		QueueBufferIngMaxMessage: config.GaoDeConfigSetting.QueueBufferIngMaxMessage,
		CompressionCodec:         config.GaoDeConfigSetting.CompressionCodec,
		Acks:                     config.GaoDeConfigSetting.Acks,
		Retries:                  config.GaoDeConfigSetting.Retries,
		RetryBackoffMs:           config.GaoDeConfigSetting.RetryBackoffMs,
		BootstrapServers:         config.GaoDeConfigSetting.BootstrapServers,
		Topic:                    config.GaoDeConfigSetting.Topic,
	})

	TencentKafkaProducer = NewKafkaProducer(&config.KafkaConfig{
		LingerMs:                 config.TencentConfigSetting.LingerMs,
		BatchSize:                config.TencentConfigSetting.BatchSize,
		QueueBuffingMaxKBytes:    config.TencentConfigSetting.QueueBuffingMaxKBytes,
		QueueBufferIngMaxMessage: config.TencentConfigSetting.QueueBufferIngMaxMessage,
		CompressionCodec:         config.TencentConfigSetting.CompressionCodec,
		Acks:                     config.TencentConfigSetting.Acks,
		Retries:                  config.TencentConfigSetting.Retries,
		RetryBackoffMs:           config.TencentConfigSetting.RetryBackoffMs,
		BootstrapServers:         config.TencentConfigSetting.BootstrapServers,
		Topic:                    config.TencentConfigSetting.Topic,
	})

	TestKafkaProducer = NewKafkaProducer(&config.KafkaConfig{
		LingerMs:                 config.TestConfigSetting.LingerMs,
		BatchSize:                config.TestConfigSetting.BatchSize,
		QueueBuffingMaxKBytes:    config.TestConfigSetting.QueueBuffingMaxKBytes,
		QueueBufferIngMaxMessage: config.TestConfigSetting.QueueBufferIngMaxMessage,
		CompressionCodec:         config.TestConfigSetting.CompressionCodec,
		Acks:                     config.TestConfigSetting.Acks,
		Retries:                  config.TestConfigSetting.Retries,
		RetryBackoffMs:           config.TestConfigSetting.RetryBackoffMs,
		BootstrapServers:         config.TestConfigSetting.BootstrapServers,
		Topic:                    config.TestConfigSetting.Topic,
	})

}

type KafkaProducer struct {
	producer *kafka.Producer // 这是基础配置

	topic      string // 这是使用配置
	groupId    string
	NotifyStop chan bool // 停止信号
}

func NewKafkaProducer(conf *config.KafkaConfig) *KafkaProducer {

	// 这里server读取配置文件
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            conf.BootstrapServers, // 其实这里是集群地址 我好像不用关心消息是发完那个机器上面的 只需要关注发给那个topic就行
		"compression.codec":            conf.CompressionCodec,
		"batch.size":                   conf.BatchSize,                // 生产者发送的消息会被分批发送，每一批的大小就是batch.size
		"queue.buffering.max.kbytes":   conf.QueueBuffingMaxKBytes,    // 生产者队列上允许的最大总消息大小总和 batch.size 只有数据累计到batch.size后，send才会发送给kafka,默认16kb
		"queue.buffering.max.messages": conf.QueueBufferIngMaxMessage, // 生产者可以缓冲的最大消息数量
		"linger.ms":                    conf.LingerMs,                 // 该值默认为5
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
		producer: producer,
		topic:    conf.Topic,
	}

	return k
}

func (k *KafkaProducer) Send(value []byte) {
	var msg *kafka.Message = nil

	// 根据规则传入的消息指定发送给对应的topic
	msg = &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny}, // 需要指定分区吗？
		Value:          value,
	}

	k.producer.Produce(msg, nil)
	time.Sleep(time.Duration(1) * time.Millisecond)
}

func (k *KafkaProducer) Close() {
	k.producer.Flush(15 * 1000)
	k.producer.Close()
}

// 3.不同消息投递不同主题
// 应该是一个topic对应一个producer
// 把需要处理的方法的对象都提出来
// 队列名有什么规范吗
