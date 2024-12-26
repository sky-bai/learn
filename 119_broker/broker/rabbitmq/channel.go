package rabbitmq

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

// rabbitChannel
type rabbitChannel struct {
	uuid       string
	connection *amqp.Connection
	channel    *amqp.Channel
	// 如果把Connection比作一条光纤电缆，那么Channel就相当于是电缆中的一束光纤
	// RabbitMQ中大部分的操作都是使用 Channel 完成的，比如：声明Queue、声明Exchange、发布消息、消费消息
}

func newRabbitChannel(conn *amqp.Connection, qos Qos) (*rabbitChannel, error) {
	rabbitCh := &rabbitChannel{
		uuid:       generateUUID(),
		connection: conn,
	}
	if err := rabbitCh.Connect(qos.PrefetchCount, qos.PrefetchSize, qos.PrefetchGlobal); err != nil {
		return nil, err
	}
	return rabbitCh, nil
}

// Connect connects to the RabbitMQ server and sets the prefetch count, size, and global
func (r *rabbitChannel) Connect(prefetchCount, prefetchSize int, prefetchGlobal bool) error {
	var err error
	r.channel, err = r.connection.Channel()
	if err != nil {
		return err
	}

	err = r.channel.Qos(prefetchCount, prefetchSize, prefetchGlobal)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitChannel) Publish(ctx context.Context, exchangeName, key string, message amqp.Publishing) error {
	if r.channel == nil {
		return errors.New("channel is nil")
	}
	return r.channel.PublishWithContext(ctx, exchangeName, key, false, false, message)
}

func (r *rabbitChannel) Close() error {
	if r.channel == nil {
		return errors.New("channel is nil")
	}
	return r.channel.Close()
}

func (r *rabbitChannel) DeclareExchange(exchangeName, kind string, durable, autoDelete bool) error {
	return r.channel.ExchangeDeclare(
		exchangeName,
		kind,
		durable,
		autoDelete,
		false,
		false,
		nil,
	)
}

// DeclareQueue 创建队列
func (r *rabbitChannel) DeclareQueue(queueName string, args amqp.Table, durable, autoDelete bool) error {
	_, err := r.channel.QueueDeclare(
		queueName,
		durable,    // durable：一个布尔值，如果为 true，则队列会在 RabbitMQ 服务器重启后仍然存在。
		autoDelete, // autoDelete：一个布尔值，如果为 true，则当最后一个消费者断开连接后，队列会被自动删除
		false,
		false,
		args,
	)
	return err
}

func (r *rabbitChannel) ConsumeQueue(queueName string, autoAck bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queueName,
		r.uuid,
		autoAck,
		false,
		false,
		false,
		nil,
	)
}

// BindQueue 将已经创建的队列绑定到一个交换器上
func (r *rabbitChannel) BindQueue(queueName, key, exchange string, args amqp.Table) error {
	return r.channel.QueueBind(
		queueName, // queueName：队列的名称。
		key,       // 路由键，用于指定消息应该路由到哪个队列。
		exchange,  // 交换器的名称。
		false,
		args,
	)
}

// 在 RabbitMQ 中，消息是发布到交换器上的，然后由交换器根据路由键将消息路由到一个或多个队列上.

// 还要写组件的测试用例
// 还要直接引用包再试一下
// 不同模式不同创建方法
