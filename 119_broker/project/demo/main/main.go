package main

import "learn/119_broker/broker/rabbitmq"

func main() {
	blurQos := rabbitmq.Qos{PrefetchCount: config.BlurService.PrefetchCount, PrefetchSize: config.BlurService.PrefetchSize, PrefetchGlobal: config.BlurService.PrefetchGlobal}
	blurExchange := rabbitmq.Exchange{Name: config.BlurService.ExchangeName, Type: config.BlurService.ExchangeType, Durable: config.BlurService.Durable}
	consumer, err := rabbitmq.NewConsumer(config.BlurService.Url, blurQos, blurExchange, config.BlurService.ConsumeQueueName, config.BlurService.ConsumeRoutingKey, blur.TestBlurImage)
}
