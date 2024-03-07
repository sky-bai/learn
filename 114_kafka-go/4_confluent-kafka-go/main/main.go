package main

import (
	"learn/114_kafka-go/4_confluent-kafka-go/producer"
	"learn/114_kafka-go/4_confluent-kafka-go/setting"
	"mosn.io/pkg/registry/dubbo/common/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		logger.Logger(.Error("init setupSetting err: %v", err))
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	producer.GaoDeKafkaProducer.Send([]byte("test"))

	// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	producer.GaoDeKafkaProducer.Close()
}
