package channel

import (
	"context"
	"learn/155_abstract/domain"
	"learn/155_abstract/provider"
)

// 渠道分发器和具体渠道的实现

// 短信渠道也要和下面的服务商做关联

type baseChannel struct {
	builder provider.SelectorBuilder
}

// smsChannel 短信渠道实现
type smsChannel struct {
	baseChannel // 这里是对渠道 渠道选择器的抽象 这里每个渠道对每个渠道选择器的实现 本身又可以抽象出一个抽象
}

func (s *smsChannel) Send(ctx context.Context, notification domain.Notification) (interface{}, error) {
	// 1.构建供应商选择器
	selector, err := s.builder.Build()
	if err != nil {
		return nil, err
	}
	// 2.选择下游
	for {
		p, err := selector.Next(ctx, notification)
		if err != nil {
			return nil, err
		}
		// 3.发送
		resp, err := p.Send(ctx, notification)
		if err == nil {
			return resp, nil
		}
	}
}

// 这里具体的实现类需要实现provider中的types里面的抽象方法 里面涉及到build 构建和next选择下游 发送send

// 作用：这里是短信渠道的抽象
// 1.实现channel里面的渠道抽象接口
// 2.里面有多个下游服务商 调用的是provider里面的types 也就是说我这里还是调用的是供应商的抽象接口
// 这里短信渠道和下游服务商通过provider/types.go里面的接口做关联

// 需求：渠道类需要管理下游多个服务商 那么我这边渠道实现类引用的是provider/types.go里面的接口

// 因为存在多个渠道管理多个服务商 所以抽象出baseChannel
