package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type Subscriber struct {
	sync.RWMutex

	ch *rabbitChannel
	c  *Consumer

	closed bool

	headers   map[string]interface{}
	queueArgs map[string]interface{}

	fn func(msg amqp.Delivery)

	queueName    string
	routingKey   string
	AutoAck      bool
	durableQueue bool
	autoDelete   bool
}

func (s *Subscriber) Unsubscribe(removeFromManager bool) error {
	s.Lock()
	defer s.Unlock()

	s.closed = true

	var err error
	if s.ch != nil {
		err = s.ch.Close()
	}

	//if s.c != nil && s.c.subscribers != nil && removeFromManager {
	//	_ = s.c.subscribers.RemoveOnly(s.routingKey)
	//}

	return err

}

func (s *Subscriber) resubscribe() {

	minResubscribeDelay := defaultMinResubscribeDelay
	maxResubscribeDelay := defaultMaxResubscribeDelay
	expFactor := defaultExpFactor
	reSubscribeDelay := defaultResubscribeDelay
	for {

		// 1.先判断自己channel是否被关闭
		closed := s.IsClosed()
		if closed {
			return
		}

		// 2.等待connection连接成功
		select {
		case <-s.c.conn.close:
			log.Error().Msg("rabbitmq consumer connection closed")
			return
		case <-s.c.conn.waitConnection:
			//log.Error().Msg("rabbitmq consumer connection waitConnection")

		}

		// 3.消费之前先判断是否connection连接成功
		s.c.mtx.Lock()
		if !s.c.conn.connected {
			s.c.mtx.Unlock()
			continue
		}

		// 4.获取到消费的实例
		channel, sub, err := s.c.conn.Consume(
			s.queueName, // 使用默认交换机 queueName = routingKey
			s.routingKey,
			s.headers,
			nil,
			s.AutoAck,
			s.durableQueue,
			s.autoDelete,
		)
		// 注意这里使用的是conn的mutex
		s.c.mtx.Unlock()

		// 5.根据err判断是否消费成功
		switch err {
		case nil:
			// 6.如果消费成功,则将channel赋值给subscriber
			reSubscribeDelay = minResubscribeDelay
			s.Lock()
			s.ch = channel
			s.Unlock()
		default:
			log.Error().Err(err).Msg("subscriber resubscribe consumer error")
			if reSubscribeDelay > maxResubscribeDelay {
				reSubscribeDelay = maxResubscribeDelay
				// 每次的延迟时间开始是100毫秒，然后每次失败后都会翻倍，直到达到30秒，然后保持在30秒，直到重新订阅
			}
			time.Sleep(reSubscribeDelay)
			reSubscribeDelay *= expFactor
			continue
			// 这里支持consumer函数异常进行重新消费的操作
			// if reSubscribeDelay > maxResubscribeDelay：这个判断语句是用来确保 reSubscribeDelay（重新订阅的延迟时间）不会超过 maxResubscribeDelay（最大的重新订阅延迟时间）。如果 reSubscribeDelay 超过了 maxResubscribeDelay，那么就将 reSubscribeDelay 设置为 maxResubscribeDelay。
			// time.Sleep(reSubscribeDelay)：这行代码会让当前的 goroutine 暂停 reSubscribeDelay 的时间。这是为了在重新订阅之前给系统一些缓冲时间，避免过于频繁的重新订阅操作。
			// reSubscribeDelay *= expFactor：这行代码会将 reSubscribeDelay 乘以一个指数因子 expFactor，这样每次重新订阅的延迟时间都会增长，直到达到 maxResubscribeDelay。这是一种常见的退避策略，用于处理可能会失败的操作，通过逐渐增加延迟时间，可以减少失败操作的频率，从而减轻系统的压力。
			// continue：这个关键字会跳过当前循环的剩余部分，直接开始下一次循环。在这里，如果重新订阅操作失败，那么就会跳过循环的剩余部分，直接开始下一次重新订阅操作。
		}

		// 7.开始消费
		//for d := range sub {
		//	//s.c.wg.Add(1)
		//	go func() {
		//		//defer s.c.wg.Done()
		//		s.fn(d)
		//	}()
		//	//s.c.wg.Done()
		//}
		for {
			select {
			case <-s.c.conn.close:
				log.Error().Msg("rabbitmq subscriber conn receive closed")
				return
			case d, ok := <-sub:
				if !ok {
					log.Error().Msg("rabbitmq subscriber channel closed")
					return
				}
				s.c.wg.Add(1)
				go func() {
					defer s.c.wg.Done()
					s.fn(d)
				}()
				//log.Info().Msg("rabbitmq subscriber ok")
			}
		}

	}
}

func (s *Subscriber) IsClosed() bool {
	s.RLock()
	defer s.RUnlock()

	return s.closed
}
