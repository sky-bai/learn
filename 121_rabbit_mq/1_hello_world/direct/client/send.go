package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// 1.
	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.DialConfig(url, amqp.Config{
		Vhost: "pre_prod",
	})
	// 一个队列允许同时消费的数量

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "test_exchange"
	routeKey := "test_route_key"

	// 交换机持久化
	// 队列持久化
	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return
	}

	q1, err := ch.QueueDeclare(
		"queue_1", // name
		false,     // durable // 是否持久化，RabbitMQ关闭后，没有持久化的Exchange将被清除
		true,      // delete when unused // 是否自动删除，如果没有与之绑定的Queue，直接删除
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	q2, err := ch.QueueDeclare(
		"queue_2", // name
		false,     // durable // 是否持久化，RabbitMQ关闭后，没有持久化的Exchange将被清除
		true,      // delete when unused // 是否自动删除，如果没有与之绑定的Queue，直接删除
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	q3, err := ch.QueueDeclare(
		"queue_3", // name
		false,     // durable // 是否持久化，RabbitMQ关闭后，没有持久化的Exchange将被清除
		true,      // delete when unused // 是否自动删除，如果没有与之绑定的Queue，直接删除
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	err = ch.QueueBind(q1.Name, routeKey, exchangeName, false, nil)
	if err != nil {
		return
	}
	err = ch.QueueBind(q2.Name, routeKey, exchangeName, false, nil)
	if err != nil {
		return
	}

	err = ch.QueueBind(q3.Name, "test_route_key_2", exchangeName, false, nil)
	if err != nil {
		return
	}

	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	for i := 0; i < 1000000; i++ {
		err = ch.PublishWithContext(ctx,
			exchangeName,       // exchange
			"test_route_key_2", // routing key
			false,              // mandatory
			false,              // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
		time.Sleep(100 * time.Millisecond)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

// 每种模式的配置文件不一样
// 1. 直连模式 需要确定的是什么 在什么场景下需要
