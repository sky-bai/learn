package domain

import "time"

// https://gitee.com/flycash/notification-platform

// Notification 通知领域模型
type Notification struct {
	ID             uint64    // 通知唯一标识
	BizID          int64     // 业务唯一标识
	Key            string    // 业务内唯一标识
	Receivers      []string  // 接收者(手机/邮箱/用户ID)
	Channel        Channel   // 发送渠道
	ScheduledSTime time.Time // 计划发送开始时间
	ScheduledETime time.Time // 计划发送结束时间
	Version        int       // 版本号
}

type Channel string

// SendResponse 发送响应
type SendResponse struct {
	NotificationID uint64     // 通知ID
	Status         SendStatus // 发送状态
}

// SendStatus 通知状态
type SendStatus string
