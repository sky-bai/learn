package rabbitmq

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var (
	DefaultAmqpConfig = amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
	}
	// Heartbeat：这是心跳间隔，设置为 10 秒。心跳是 RabbitMQ 用来检测客户端是否仍然在线的机制。如果在这个间隔内没有收到客户端的任何消息，RabbitMQ 会认为客户端已经断开连接。
	// Locale：这是本地化设置，设置为 "en_US"。这个设置通常用于错误消息和其他从服务器返回的文本的本地化。
)

type Qos struct {
	PrefetchCount  int
	PrefetchSize   int
	PrefetchGlobal bool
}

type Exchange struct {
	Name    string
	Type    string // "direct", "fanout", "topic", "headers"
	Durable bool
}

type rabbitConnection struct {
	sync.Mutex
	Connection *amqp.Connection
	Channel    *rabbitChannel // 用于创建exchange,声明Queue、声明Exchange、发布消息、消费消息

	url      string
	exchange Exchange
	qos      Qos

	connected      bool
	close          chan bool
	waitConnection chan struct{}
}

func newRabbitMQConnection(url string, qos Qos, exchange Exchange) *rabbitConnection {
	conn := &rabbitConnection{
		url:            url,
		qos:            qos,
		exchange:       exchange,
		close:          make(chan bool),
		waitConnection: make(chan struct{}),
	}

	close(conn.waitConnection)
	// 因为这里close了这个channel 当后面处理消息的时候会读这个已关闭的channel

	return conn

}

// Connect 判断是否已经连接，如果没有连接则进行连接
func (r *rabbitConnection) Connect(secure bool, config *amqp.Config) error {
	r.Lock()

	if r.connected {
		r.Unlock()
		return nil
	}

	select {
	case <-r.close:
		r.close = make(chan bool)
	default:
	}

	r.Unlock()

	return r.connect(secure, config)
}

// connect 执行tryConnect和reconnect
func (r *rabbitConnection) connect(secure bool, config *amqp.Config) error {
	// 1.建立连接
	if err := r.tryConnect(secure, config); err != nil {
		return err
	}
	//log.Info().Msg("mq connect ok")

	r.Lock()
	r.connected = true
	r.Unlock()

	// 2.开启携程等待异常消息
	go r.reconnect(secure, config)
	return nil
}

// reconnect 用于重连
func (r *rabbitConnection) reconnect(secure bool, config *amqp.Config) {
	ctx := context.Background()
	var connect bool
	// 一直循环等待异常消息

	for {
		if connect {
			if err := r.tryConnect(secure, config); err != nil {
				time.Sleep(2 * time.Second)
				continue
			}

			r.Lock()
			r.connected = true
			r.Unlock()
			close(r.waitConnection)
		}
		connect = true

		// 1.获取rabbitmq服务器返回的错误信息
		notifyClose := make(chan *amqp.Error)
		r.Connection.NotifyClose(notifyClose)

		// 2.获取rabbitmq服务器返回的消息
		channelNotifyReturn := make(chan amqp.Return)
		r.Channel.channel.NotifyReturn(channelNotifyReturn)

		select {
		// 接受退出信号
		case result, ok := <-channelNotifyReturn:
			if !ok {
				log.Error().Msg("rabbit connection channelNotifyReturn channel closed")
				// Channel closed, probably also the channel or connection.
				return
			}
			// when a Publishing is unable to be delivered either due to the `mandatory` flag set and no route found, or `immediate` flag set and no free consumer
			log.Ctx(ctx).Error().Msgf("reconnect notify error reason: %s, description: %s", result.ReplyText, result.Exchange)
		case err := <-notifyClose:
			log.Ctx(ctx).Error().Msgf("reconnect notifyClose: %v", err)
			r.Lock()
			r.connected = false
			r.waitConnection = make(chan struct{})
			r.Unlock()
		case <-r.close:
			return
		}
	}

}

// tryConnect 用于真正连接
func (r *rabbitConnection) tryConnect(secure bool, config *amqp.Config) error {
	// 1.设置dial配置
	if config == nil {
		config = &DefaultAmqpConfig
	}

	if r.connected {
		log.Info().Msg("close old connection")
		err := r.Connection.Close()
		if err != nil {
			log.Error().Err(err).Msg("close old connection err")
		} else {
			log.Info().Msg("close old connection ok")
		}
	}

	// 2.dial连接
	var err error
	r.Connection, err = amqp.DialConfig(r.url, *config)
	if err != nil {
		return err
	}

	// 3.创建connection和channel
	if r.Channel, err = newRabbitChannel(r.Connection, r.qos); err != nil {
		return err
	}

	// 4.声明exchange
	if r.exchange.Name != "" {
		_ = r.Channel.DeclareExchange(r.exchange.Name, r.exchange.Type, r.exchange.Durable, false)
	}

	return err
}

// Close 关闭连接
func (r *rabbitConnection) Close() error {
	r.Lock()
	defer r.Unlock()

	select {
	case <-r.close:
		return nil
	default:
		close(r.close)
		r.connected = false
	}

	if r.Connection == nil {
		return errors.New("connection is nil")
	}

	return r.Connection.Close()
}

func (r *rabbitConnection) DeclarePublishQueue(queueName, routingKey string, bindArgs, queueArgs amqp.Table, durableQueue, autoDel bool) error {

	// 1.声明要生产的 queue
	if err := r.Channel.DeclareQueue(queueName, queueArgs, durableQueue, autoDel); err != nil {
		return err
	}

	// 2.绑定queue
	if err := r.Channel.BindQueue(queueName, routingKey, r.exchange.Name, bindArgs); err != nil {
		return err
	}

	return nil
	// 这里先调用 ConsumeQueue 再调用 BindQueue 的原因可能是为了确保在绑定操作完成之前，消费者已经开始监听队列中的消息。
	// 这样，一旦队列与交换器绑定成功，消费者就可以立即开始接收到通过交换器路由过来的消息。
	// 如果先进行绑定操作，然后再开始消费消息，那么在绑定操作和消费操作之间的这段时间里，如果有消息被发布到交换器，那么这些消息可能会丢失（因为此时还没有消费者开始监听队列）

}

func (r *rabbitConnection) Publish(ctx context.Context, exchangeName, routingKey string, msg amqp.Publishing) error {
	return r.Channel.Publish(ctx, exchangeName, routingKey, msg)
}

func (r *rabbitConnection) Consume(queueName, routingKey string, bindArgs, qArgs amqp.Table, autoAck, durableQueue, autoDel bool) (*rabbitChannel, <-chan amqp.Delivery, error) {

	if err := r.Channel.DeclareQueue(queueName, qArgs, durableQueue, autoDel); err != nil {
		return nil, nil, err
	}

	deliveries, err := r.Channel.ConsumeQueue(queueName, autoAck)
	if err != nil {
		return nil, nil, err
	}

	if r.exchange.Name != "" {
		if err = r.Channel.BindQueue(queueName, routingKey, r.exchange.Name, bindArgs); err != nil {
			return nil, nil, err
		}
	}

	return r.Channel, deliveries, nil
}
