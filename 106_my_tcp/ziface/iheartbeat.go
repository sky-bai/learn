package ziface

import "time"

type IHeartbeatChecker interface {
	SetOnRemoteNotAlive(OnRemoteNotAlive)
	Set(conn IConnection, interval, afterTime time.Duration) // 设置当前链接在afterTime时间后启动心跳检测
	Stop(conn IConnection)                                   // 停止当前链接的心跳检测
	ResetConnExpire(conn IConnection, addTime time.Duration) // 重置当前链接的心跳检测 根据当前时间将当前链接的心跳检测时间更新为interval之后 setTimer方法将会被取消
	BindConnManger(manager IConnManager)
	Clone() IHeartbeatChecker
}

// OnRemoteNotAlive // 因为这里通过心跳检测器去管理链接了，所有在这一层提供用户自定义的远程连接不存活时的处理方法
type OnRemoteNotAlive func(IConnection)

type HeartBeatOption struct {
	OnRemoteNotAlive OnRemoteNotAlive //用户自定义的远程连接不存活时的处理方法
}
