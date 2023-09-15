package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
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

	exchangeName := "work_queues_exchange"
	routeKey := "work_queues_route_key"

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
		"work_queue_queue", // name
		false,              // durable // 是否持久化，RabbitMQ关闭后，没有持久化的Exchange将被清除
		true,               // delete when unused // 是否自动删除，如果没有与之绑定的Queue，直接删除
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)

	err = ch.QueueBind(q1.Name, routeKey, exchangeName, false, nil)
	if err != nil {
		return
	}

	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := bodyFrom(os.Args)
	//body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		exchangeName,            // exchange
		"work_queues_route_key", // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	select {}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
