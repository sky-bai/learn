package rabbitmq

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type Producer struct {
	conn *rabbitConnection
}

// 这里结构体只需要实现三方的接口即可 组装结构体已经实现好了的方法

func NewProducer(url string, qos Qos, exchange Exchange) (*Producer, error) {

	// 1.这里初始化的时候就初始化好
	r := &Producer{
		conn: newRabbitMQConnection(url, qos, exchange),
	}
	// 2.启动的时候建立连接
	err := r.conn.Connect(false, &DefaultAmqpConfig)
	if err != nil {
		log.Error().Err(err).Msg("rabbitmq broker connect err")
		return nil, err
	}

	return r, nil
}

func (p *Producer) Publish(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	return p.publish(ctx, routingKey, msg)
}

func (p *Producer) publish(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	if p.conn == nil {
		return errors.New("connection is nil")
	}
	return p.conn.Publish(ctx, p.conn.exchange.Name, routingKey, msg)
}

func (p *Producer) Close() error {
	if p.conn == nil {
		log.Error().Msg("Close rabbitmq broker conn is nil")
		return nil
	}

	return p.conn.Close()
}
