package main

import (
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
}

func setupConfig() error {
	setting, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = setting.ReadSection("GaoDe", &config.GaoDeConfigSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Tencent", &config.TencentConfigSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Test", &config.TestConfigSetting)
	if err != nil {
		return err
	}

	data, _ := json.Marshal(config.TestConfigSetting)
	fmt.Println("----", string(data))

	return nil
}

func main() {
	for i := 0; i < 10; i++ {
		producer.TestKafkaProducer.Send([]byte("test"))
		//producer.GaoDeKafkaProducer.Send([]byte("gaoDe test"))
		time.Sleep(1 * time.Second)
	}
	//
	//producer.TencentKafkaProducer.Send([]byte("tencent test"))
	//
	//// 等待中断信号
	quit := make(chan os.Signal)

	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//
	//producer.GaoDeKafkaProducer.Close()
	//producer.TencentKafkaProducer.Close()
	producer.TestKafkaProducer.Close()
}
