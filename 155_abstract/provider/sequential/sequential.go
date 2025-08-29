package sequential

import (
	"context"
	"fmt"
	"learn/155_abstract/domain"
	"learn/155_abstract/errs"
	"learn/155_abstract/provider"
)

// 线性轮训

// 这里的选择服务商策略是顺序选择的原因是适用于服务商相同的情况下，或者说没有制定服务商的情况，比如短信渠道
// 如果是apple push这种渠道，可能会有多个服务商，但是每个服务商的能力和覆盖范围不一样，这种情况下就不适合顺序选择就应该根据实际参数选择

// 这里是可以管理多个供应商的选择策略

// 这里实现provider.Selector接口 通过策略管理多个供应商

// 渠道下面是供应商选择器
// 本层就是供应商选择器的实现

// 编译时接口检查技巧

// selector 供应商顺序选择器 就需要思考自己选择多个供应商的逻辑
type selector struct {
	providers []provider.Provider // 1.先保存多个供应商 这里抽象出来的接口
	idx       int                 // 2.再定义一个索引 用于记录当前选择到哪个供应商了
}

// Next 这里就是选择下一个供应商的逻辑
func (s *selector) Next(_ context.Context, _ domain.Notification) (provider.Provider, error) {
	if len(s.providers) == 0 {
		return nil, fmt.Errorf("%w", errs.ErrNoAvailableProvider)
	}

	p := s.providers[s.idx%len(s.providers)]
	return p, nil
}

// SelectorBuilder 供应商选择器的构造器 需要保存多个供应商
type SelectorBuilder struct {
	providers []provider.Provider
}

func NewSelectorBuilder(providers []provider.Provider) *SelectorBuilder {
	return &SelectorBuilder{providers: providers}
}

func (s *SelectorBuilder) Build() (provider.Selector, error) {
	return &selector{
		providers: s.providers,
	}, nil
}
