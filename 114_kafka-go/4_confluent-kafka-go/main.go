package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learn/114_kafka-go/4_confluent-kafka-go/config"
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

	producer.TestKafkaProducer = producer.NewKafkaProducer(&config.KafkaConfig{
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

	producer.TransactionProducer = producer.NewKafkaProducer(&config.KafkaConfig{
		BootstrapServers: config.TransactionSetting.BootstrapServers,
		Topic:            config.TransactionSetting.Topic,
		Partition:        config.TransactionSetting.Partition,
		TransactionId:    config.TransactionSetting.TransactionId,
		LingerMs:         config.TransactionSetting.LingerMs,
		BatchSize:        config.TransactionSetting.BatchSize,
		CompressionCodec: config.TransactionSetting.CompressionCodec,
		Acks:             config.TransactionSetting.Acks,
		Retries:          config.TransactionSetting.Retries,
	})
}

func setupConfig() error {
	setting, err := config.NewConfig()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Transaction", &config.TransactionSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Test", &config.TestConfigSetting)
	if err != nil {
		return err
	}

	data, _ := json.Marshal(config.TransactionSetting)
	fmt.Println("----", string(data))

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
	//for i := 0; i < 50; i++ {
	//	producer.TestKafkaProducer.Send([]byte("wxy"), CustomPartition)
	//	time.Sleep(1 * time.Second)
	//	fmt.Printf("i %d\n", i)
	//}

	TransactionUsage()

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// close
	producer.TransactionProducer.Close()
	producer.TestKafkaProducer.Close()
}

func TransactionUsage() {

	maxDuration, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatalf("time.ParseDuration err: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), maxDuration)
	defer cancel()

	err = producer.TransactionProducer.InitTransactions(ctx)
	if err != nil {
		fmt.Println("errrr:", err)
		return
	}

	for i := 0; i < 10; i++ {
		producer.TransactionProducer.BeginTransaction()
		err := producer.TransactionProducer.Send([]byte("wxy"), nil)
		if err != nil {
			fmt.Println("err:", err)
			producer.TransactionProducer.AbortTransaction(ctx)
			return
		}
		time.Sleep(time.Second)
	}

	producer.TransactionProducer.CommitTransaction(ctx)
	fmt.Println("done")

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
