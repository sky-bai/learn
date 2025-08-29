package provider

import (
	"context"
	"learn/155_abstract/domain"
)

// provider 这一层 就是实际的实现类 比如短信 邮件
// 那么这一层的方法就是发送方法 实现类要发送
// 然后定义该层的接口就是provider接口里面有发送方法

type Provider interface {
	Send(ctx context.Context, notification domain.Notification) (domain.SendResponse, error)
}

// 这里定义实现类的抽象方法 需要按照什么标准来实现规范
// 这一层相当于定义每一个渠道的实现类的规范是什么
// 这里定义两个 一个是实现类的构造方法 一个是该实现类的发送方法
// 使用方法就是 1.先build 2.再next选择下游 3.最后send发送

// 这里涉及到多个服务商 需要如何选择下游

// Selector 供应商选择器接口
type Selector interface {
	// Next 获取下一个供应商，无可用供应商时返回错误
	Next(ctx context.Context, notification domain.Notification) (Provider, error)
	// next 返回的其实是一个provider 实现类实例
}

// SelectorBuilder 供应商选择器的构造器
type SelectorBuilder interface {
	// Build 构造选择器，可以在Build方法上添加参数来构建更复杂的选择器
	Build() (Selector, error)
	// build 返回的其实是一个selector 实现类实例
}

// 先build出一个selector 然后再selector出一个provider
// 因为这里是一个渠道有多个供应商 渠道和多个供应商之前的关系
// 渠道用SelectorBuilder 供应商实现Provider

// 作用：定义每个渠道实现类的规范 注意是渠道不是下面的供应商

// Selector：表示“行为策略”，它定义了如何选择一个 Provider 的规则（例如轮询、随机、权重优先等）。
// SelectorBuilder：表示“构建策略的工厂”，它封装了如何创建对应的 Selector。
// 思考方向：
// 业务中是否存在“多种可替换的策略”？
// 是否有必要通过一个“工厂接口（Builder）”动态选择不同的策略？

// Selector
//
//职责：封装“选择提供商”的逻辑。
// SelectorBuilder
//
//职责：负责“如何创建不同类型的 Selector 实例”。

// todo 上层是如何使用下一层的
