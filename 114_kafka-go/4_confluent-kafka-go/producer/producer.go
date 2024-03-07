package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"time"
)

// 什么时候会出现多个一个消费者组消费多个topic 动态更新缓冲区大小

// 在上层做topic的区分

// 不同topic有不同的生产者配置

var GaoDeKafkaProducer *KafkaProducer

func init() {
	var GaoDe KafkaConfig //
	GaoDeKafkaProducer = NewKafkaProducer(GaoDe)

}

type KafkaProducer struct {
	producer *kafka.Producer // 这是基础配置

	topic      string // 这是使用配置
	groupId    string
	NotifyStop chan bool // 停止信号
}

type KafkaConfig struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`

	GroupId                   string `json:"group.id"`
	BootstrapServers          string `json:"bootstrap.servers"`
	CompressionCodec          string `json:"compression.codec"`
	QueueBufferingMaxKBytes   int    `json:"queue.buffering.max.kbytes"`
	QueueBufferingMaxMessages int    `json:"queue.buffering.max.messages"`
	LingerMs                  int    `json:"linger.ms"`
}

func NewKafkaProducer(conf KafkaConfig) *KafkaProducer {

	// 这里server读取配置文件
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            conf.BootstrapServers, // 其实这里是集群地址 我好像不用关心消息是发完那个机器上面的 只需要关注发给那个topic就行
		"compression.codec":            conf.CompressionCodec,
		"queue.buffering.max.kbytes":   conf.QueueBufferingMaxKBytes,   // 生产者队列上允许的最大总消息大小总和 batch.size 只有数据累计到batch.size后，send才会发送给kafka,默认16kb
		"queue.buffering.max.messages": conf.QueueBufferingMaxMessages, // 生产者可以缓冲的最大消息数量
		"linger.ms":                    conf.LingerMs,                  // 该值默认为5
	})

	if err != nil {
		panic(err)
	}
	defer producer.Close()

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

// 不同消息投递不同主题
// 应该是一个topic对应一个producer

// 把需要处理的方法的对象都提出来
