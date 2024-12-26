package rabbitmq

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"sync"
)

type ConsumerHandler func(ctx context.Context, msg amqp.Delivery) error

type Consumer struct {
	mtx        sync.Mutex
	conn       *rabbitConnection
	QueueName  string
	routingKey string
	handler    ConsumerHandler

	wg sync.WaitGroup // 并发消费

	//subscribers *broker.SubscriberSyncMap
}

func NewConsumer(url string, qos Qos, exchange Exchange, QueueName, routingKey string, handler ConsumerHandler) (*Consumer, error) {
	c := &Consumer{
		//subscribers: broker.NewSubscriberSyncMap(),
		routingKey: routingKey,
		handler:    handler,
		QueueName:  QueueName,
	}

	c.conn = newRabbitMQConnection(url, qos, exchange)
	conf := DefaultAmqpConfig

	err := c.conn.Connect(false, &conf)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (b *Consumer) SetHandler(handler ConsumerHandler) {
	b.handler = handler
}

func (b *Consumer) Disconnect() error {
	if b.conn == nil {
		return errors.New("connection is nil")
	}
	b.wg.Wait()

	//b.subscribers.Clear()

	ret := b.conn.Close()

	return ret
}

func (b *Consumer) Subscribe(routingKey string) (*Subscriber, error) {
	if b.conn == nil {
		log.Error().Msgf("connection is nil")
		return nil, errors.New("connection is nil")
	}

	// 1.处理消息是否ack的逻辑
	fn := func(msg amqp.Delivery) {

		err := b.handler(context.Background(), msg)
		if err == nil {
			_ = msg.Ack(false)
		} else if err != nil {
			_ = msg.Nack(false, true)
		}

	}

	// 2.创建subscriber
	sub := &Subscriber{
		queueName:    b.QueueName,
		routingKey:   routingKey,
		c:            b,
		durableQueue: true,
		autoDelete:   false,
		fn:           fn,
		headers:      nil,
		queueArgs:    nil,
	}

	// 3.添加到消费者的订阅者列表
	//b.subscribers.Add(routingKey, sub)

	//go sub.resubscribe()

	return sub, nil

}

func (b *Consumer) Serve() {

	sub, err := b.Subscribe(b.routingKey)
	if err != nil {
		log.Fatal().Err(err).Msg("rabbitmq consumer subscribe error")
	}

	sub.resubscribe()

}

func (b *Consumer) Stop() {
	err := b.Disconnect()
	if err != nil {
		log.Error().Err(err).Msg("rabbitmq consumer disconnect error")
	}
}

// 为了让生产者和消费者充分的解耦, 理想情况下, 生产者仅仅知道关于交换机的信息(而不是队列), 并且消费者仅仅知道关于队列的信息(而不是交换机). 绑定关系表明交换机和队列的关系
//
//一种可能能方式是让生产者处理交换机的创建, 消费者创建队列并将队列绑定在交换机上. 这种去耦合的方式的优点是, 如果消费者需要队列, 仅仅是按照需求, 简单的创建队列并绑定它们就可以了, 并不需要让生产者知道任何关于队列的信息. 但这并不是充分的解耦: 因为消费者为了绑定必须要知道这个交换机.
//
//2. 生产者创建一切
//在生产者运行的时候, 可以配置去创建所有需要的组件(交换机, 队列和绑定关系). 这种方式的好处是, 没有消息会丢失掉(因为队列已经创建好并绑定到交换机上了, 并不需要让任何消费者先启动).
//
//然而, 这就意味着, 生产者必须要知道所有需要和交换机绑定的队列. 这是一种耦合度很高的方式. 原因是每次需要添加新队列的时候, 生产者必须要重新配置和部署以用来创建队列和绑定队列
//
//3. 消费者创建一切
//相反的方式是, 在消费者运行的时候, 让消费者去创建它需要的交换机, 队列和绑定关系. 和前一种方式一样, 这种方式产生了耦合, 因为消费者必须要知道它们要绑定队列的交换机的信息. 交换机如果有了任何改动(例如重命名), 意味着所有的消费者必须要重新配置和部署. 当存在大量队列和消费者时, 这种复杂性可能会令人望而却步.
