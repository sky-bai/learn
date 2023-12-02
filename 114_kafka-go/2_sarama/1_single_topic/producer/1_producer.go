package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

// 创建一个分区

func main() {

	// 创建主题通过什么去创建

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

	//topic := "nil-topic"
	topic := "two"
	brokers := []string{"47.106.250.122:9092", "47.119.157.148:9092", "47.112.177.81:9092"}

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		fmt.Println("kafka Failed to start consumer: ", err)
		return
	}

	// 如果生产者往一个没有创建的topic里面丢信息会发生什么？

	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("test"),
	}

	err = producer.Close()
	if err != nil {
		fmt.Println("kafka Failed to close producer: ", err)
		return
	}

	fmt.Println("Down!")

}

// 如何查看kafka的错误日志

// 0:生产者发送过来的数据，不需要等数据罗盘应答
// 1:生产者发送过来的数据，需要等待leader应答，不需要等待follower应答
// -1:生产者发送过来的数据，需要等待leader应答，需要等待follower应答

// 分区进行数据切割
// 没有指定分区 直接往topic写入数据，则找默认的分区
// key的hash值 % 分区数 = 分区id

// 1.确定消费者顺序
// 2.确定消息何时删除
// 3.确定消费组
// 4.kafka 商家视频
// 5.回过来通过UI查看分区数据

// 如何把多张表的数据发送给单个分区
// 把表名放入key中然后就发送给单个分区

// 自定义分区器
// 1.实现kafka.Partitioner接口
// 对于某个消息

// GpsSend0 根据消息内容发往指定分区
func GpsSend0() {

}

// 生产者如何提高吞吐量
// 批量拉还是一个一个拉
// batch.size 批次大小 默认是16k
// linger.ms 等待时间 默认是0ms 修改 5-100ms 来一个就拉一个
// 修改5-10ms 会有一定的提升
// compression.type 压缩类型 snappy 默认是none
// RecordAccumulator 缓冲区 默认是32M 修改为64M 这个是根据什么来修改的昵

// producer 的事务id是唯一的 重启后又能处理未完成的事务

// sarama使用事务的列子
// 只使用消费者组的例子
