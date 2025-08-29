package channel

import (
	"context"
	"fmt"
	"learn/155_abstract/domain"
)

// 渠道抽象+渠道分发策略+分发器实现+具体实现

type Channel interface {
	Send(ctx context.Context, notification domain.Notification) (interface{}, error)
}

// 接口是上层调用以及自我内部规范 上级调用是发送策略类调用它，需要讲具体的实现注入到发送策略类中
// 渠道的发送接口 是根据通知信息来做处理
// 写完接口 写实现类 实现类返回接口
// 渠道的实现应该是多个渠道(map 数组)，然后可以分发给多个渠道
// 这里的分发器其实是管理多个渠道，选择好渠道之后还是调用自己的Send发送接口

// 这一层渠道是不同的渠道，每个渠道没有关系 这里是对多个渠道的抽象

type Dispatcher struct {
	channels map[domain.Channel]Channel
}

func (d *Dispatcher) Send(ctx context.Context, notification domain.Notification) (interface{}, error) {
	ch, ok := d.channels[notification.Channel]
	if !ok {
		return nil, fmt.Errorf("no available channel")
	}
	return ch.Send(ctx, notification)
}

func NewDispatcher(channels map[domain.Channel]Channel) *Dispatcher {
	return &Dispatcher{
		channels: channels,
	}
}

// 这里的上下层联系可以用接口 和 实现类

// 作用：管理多个渠道的抽象
