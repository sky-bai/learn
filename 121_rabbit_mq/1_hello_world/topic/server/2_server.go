package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError2(err error, msg string) {
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
	failOnError2(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"exchange_mode_topic", // name
		"topic",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println("ExchangeDeclare", err)
		return
	}

	queue_2, err := ch.QueueDeclare(
		"queue_2", // name
		false,     // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError2(err, "Failed to declare a queue")

	msgs_2, err := ch.Consume(
		queue_2.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError2(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs_2 {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
