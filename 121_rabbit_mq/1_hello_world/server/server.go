package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	url := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.DialConfig(url, amqp.Config{
		Vhost: "pre_prod",
	})
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"test_exchange", // name
		"direct",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println("ExchangeDeclare", err)
		return
	}

	queue_1, err := ch.QueueDeclare(
		"queue_1", // name
		false,     // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Println(queue_1.Name)

	msgs, err := ch.Consume(
		queue_1.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
