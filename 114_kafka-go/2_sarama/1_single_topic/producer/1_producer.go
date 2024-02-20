package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

// 创建一个分区
var kafkaBrokers = []string{"47.119.157.148:9092"}
var topics = "test-z"

var (
	wg                                   sync.WaitGroup
	enqueued, timeout, successes, errors int
)

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

	// 异步生产者不建议把 Errors 和 Successes 都开启，一般开启 Errors 就行
	// 同步生产者就必须都开启，因为会同步返回发送成功或者失败

	producer, err := sarama.NewAsyncProducer(kafkaBrokers, config)
	if err != nil {
		fmt.Println("kafka Failed to start consumer: ", err)
		return
	}

	// todo 如果生产者往一个没有创建的topic里面丢信息会发生什么？ 会创建一个只有一个分区的topic 但是如何查看这个分区在那一台机器上面昵
	limit := 10
	for i := 0; i < limit; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topics,
			Value: sarama.StringEncoder("test"),
		}
		fmt.Println("111")

		// 异步发送只是写入内存了就返回了，并没有真正发送出去
		// sarama 库中用的是一个 channel 来接收，后台 goroutine 异步从该 channel 中取出消息并真正发送
		// select + ctx 做超时控制,防止阻塞 producer.Input() <- msg 也可能会阻塞
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
		select {
		case producer.Input() <- msg:
			enqueued++
		case <-ctx.Done():
			timeout++
		}
		cancel()
		if i%10000 == 0 && i != 0 {
			log.Printf("已发送消息数:%d 超时数:%d\n", i, timeout)
		}
	}

	// todo 如何查看发送的数据是否问题

	err = producer.Close()
	if err != nil {
		fmt.Println("kafka Failed to close producer: ", err)
		return
	}

	fmt.Println("Down!")
	wg.Wait()

	log.Printf("发送完毕 总发送条数:%d enqueued:%d timeout:%d successes: %d errors: %d\n", limit, enqueued, timeout, successes, errors)

}

// auto.create.topics.enable：是否允许自动创建 Topic 如何查看这个参数 生产环境一般不设置自动

// todo 删除没有用的kafka topic 有什么影响

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
