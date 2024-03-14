package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
	"learn/114_kafka-go/4_confluent-kafka-go/consumer"
	"learn/114_kafka-go/4_confluent-kafka-go/producer"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := setupConfig()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	//producer.TransactionProducer = producer.NewKafkaProducer(&config.KafkaConfig{
	//	BootstrapServers: config.ChannelKafkaConfigSetting.Transaction.BootstrapServers,
	//	Topic:            config.ChannelKafkaConfigSetting.Transaction.Topic,
	//	Partition:        config.ChannelKafkaConfigSetting.Transaction.Partition,
	//	TransactionId:    config.ChannelKafkaConfigSetting.Transaction.TransactionId,
	//	LingerMs:         config.ChannelKafkaConfigSetting.Transaction.LingerMs,
	//	BatchSize:        config.ChannelKafkaConfigSetting.Transaction.BatchSize,
	//	CompressionCodec: config.ChannelKafkaConfigSetting.Transaction.CompressionCodec,
	//	Acks:             config.ChannelKafkaConfigSetting.Transaction.Acks,
	//	Retries:          config.ChannelKafkaConfigSetting.Transaction.Retries,
	//})

	producer.TestKafkaProducer = producer.NewKafkaProducer(&config.KafkaConfig{
		LingerMs:                 config.ChannelKafkaConfigSetting.Test.LingerMs,
		BatchSize:                config.ChannelKafkaConfigSetting.Test.BatchSize,
		QueueBuffingMaxKBytes:    config.ChannelKafkaConfigSetting.Test.QueueBuffingMaxKBytes,
		QueueBufferIngMaxMessage: config.ChannelKafkaConfigSetting.Test.QueueBufferIngMaxMessage,
		CompressionCodec:         config.ChannelKafkaConfigSetting.Test.CompressionCodec,
		Acks:                     config.ChannelKafkaConfigSetting.Test.Acks,
		Retries:                  config.ChannelKafkaConfigSetting.Test.Retries,
		RetryBackoffMs:           config.ChannelKafkaConfigSetting.Test.RetryBackoffMs,
		BootstrapServers:         config.ChannelKafkaConfigSetting.Test.BootstrapServers,
		Topic:                    config.ChannelKafkaConfigSetting.Test.Topic,
	})

	consumer.TestKafkaConsumer = consumer.NewKafkaConsumer(&config.KafkaConsumerConfig{
		BootstrapServers: config.ChannelKafkaConfigSetting.Consumer.BootstrapServers,
		Topic:            config.ChannelKafkaConfigSetting.Consumer.Topic,
		GroupId:          config.ChannelKafkaConfigSetting.Consumer.GroupId,
		AutoOffsetReset:  config.ChannelKafkaConfigSetting.Consumer.AutoOffsetReset,
	})

}

func setupConfig() error {
	setting, err := config.NewConfig()
	if err != nil {
		return err
	}

	err = setting.ReadSection(config.ChannelKafkaConfigSetting)
	if err != nil {
		log.Fatalf("vp.Unmarshal err: %v", err)
	}

	data, _ := json.Marshal(config.ChannelKafkaConfigSetting)
	fmt.Println("----setupConfig ", string(data))

	return nil
}

// 比如一个订单表，会把一个表名字传入过来 然后这一张表的数据都在同一个分区上面
// CustomPartition 自定义分区器

func CustomPartition(value string) (partition int32) {
	if value == "wxy" {
		return 1
	}
	return 0
}

func main() {
	doneChannel := make(chan struct{})

	go consumerNew()

	//time.Sleep(1 * time.Second)
	go producerNew()

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// close
	producer.TestKafkaProducer.Close()

	consumer.TestKafkaConsumer.Close()
}

func producerNew(doneChannel chan struct{}) {

	for {
		select {
		case <-doneChannel:
			fmt.Println("生产者退出")
			return
		default:
			err := producer.TestKafkaProducer.Send([]byte("wxy"), CustomPartition)
			if err != nil {
				fmt.Println("Send err:", err)
			}
			time.Sleep(1 * time.Second)
			fmt.Printf(" producer i %d\n", i)

		}

	}

	// 先关闭生产者
	// 再关闭消费者
	// 再关闭客户端
}

// 单个消费者 还要多个消费者

func consumerNew(doneChannel chan struct{}) {

	for {
		select {
		case <-doneChannel:
			fmt.Println("消费者退出")
			return
		default:
			consumer.TestKafkaConsumer.Consumer(consumer.HandlerTest)
		}

	}

}

func TransactionUsage() {
	maxDuration, err := time.ParseDuration("15s")
	if err != nil {
		log.Fatalf("time.ParseDuration err: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), maxDuration)
	defer cancel()

	err = producer.TransactionProducer.InitTransactions(ctx)
	if err != nil {
		fmt.Println("InitTransactions err:", err)
		return
	}

	for i := 0; i < 10; i++ {

		err = producer.TransactionProducer.BeginTransaction()
		if err != nil {
			fmt.Println("BeginTransaction err:", err)
			return
		}
		err = producer.TransactionProducer.Send([]byte("wxy"), nil)
		if err != nil {
			fmt.Println("Send err:", err)
			producer.TransactionProducer.AbortTransaction(nil)
			return
		}
		time.Sleep(time.Second)

		err = producer.TransactionProducer.CommitTransaction(nil)
		if err != nil {
			fmt.Println("CommitTransaction err:", err)
			err = producer.TransactionProducer.AbortTransaction(nil)
			if err != nil {
				fmt.Println("AbortTransaction err:", err)
			}

			continue
			return
		}
		fmt.Println("done")
	}

}

// kafka概念：
//主题 业务的分类，放什么类型的数据
//分区 主题可以分多个分区 物理上对数据分片
//副本 等同分区，多副本下只有主副本（分区）对生产者消费者提供服务
//分区和副本分布在不同节点
//只有一个节点，所有分区就在一个节点上
//只有一个节点多副本也没有意义，所以多节点副本要大于1，小于节点数
//一个分区就是一个队列
//生产者 ：
//生产者生产主题消息时，会根据策略分发到不同分区
//消费者：
//消费者有消费者组概念
//消费者组对应的就是主题，这个组要消费哪个主题的数据
//两个组消费同一个主题数据，主题数据就会消费两遍
//每个组内成员消费分区数据时都会记录消费到哪里了（偏移量）
//所以说同一组内成员不能消费同一个分区数据，会重复消费（偏移量不共享）
//组内消费者大于分区数量，多出的消费者会空闲（为了不重复消费）
//反之分区大于组内消费者，就会有消费者消费多个分区数据

// 如何在指定机器创建分区
