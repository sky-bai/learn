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

	exchangeName := "publish_sub_logs"
	//routeKey := "route_key_logs"

	// 交换机持久化
	// 队列持久化
	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout", // 对于fanout类型的交换机 routeKey无效
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return
	}

	failOnError(err, "Failed to declare a queue")

	// 2.创建一个具有生成名称的非持久队列
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	// 3.绑定队列
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := bodyFrom(os.Args)
	//body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
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
