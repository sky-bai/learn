package ziface

import (
	"time"
)

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器方法
	Start()
	// Stop 停止服务器方法
	Stop()
	// Serve 开启业务服务方法
	Serve()

	// AddHandler 给当前服务注册一个handler业务方法，供客户端链接处理使用
	AddHandler(msgType string, handler func(request IRequest))

	SetOnConnStop(func(IConnection))  // SetOnConnStop 设置该Server的连接断开时的Hook函数
	GetOnConnStop() func(IConnection) //得到该Server的连接断开时的Hook函数

	StartHeartBeatWithOption(time.Duration, *HeartBeatOption) //启动心跳检测(自定义回调)

	GetConnMgr() IConnManager               // GetConnMgr 得到链接管理
	GetMsgHandler() IMsgHandle              //获取Server绑定的消息处理模块
	GetHeartBeatChecker() IHeartbeatChecker //获取心跳检测器

}
