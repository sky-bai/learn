package client

import (
	"context"
	"learn/155_abstract/domain"
)

// 因为这里有多个服务商的SDK，所以把每个服务商的SDK都抽象成一个接口
// 每个客户端就有一个发送方法就可以了

type Client interface {
	Send(ctx context.Context, notification domain.Notification) (domain.SendResponse, error)
}
