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
