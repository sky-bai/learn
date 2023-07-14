package ziface

import (
	"context"
	"net"
	"time"
)

/*
 IConnection 定义连接接⼝
*/

type IConnection interface {
	Start()                          // Start 启动连接，让当前连接开始⼯工作
	Stop()                           // Stop 停⽌止连接，结束当前连接状态M
	Write(data string) error         // 直接将Message数据发送给远程的TCP客户端
	GetTCPConnection() *net.TCPConn  // 获取当前连接绑定的socket conn
	GetHeartBeat() IHeartbeatChecker // 获取当前连接绑定的心跳检测器
	GetMsgHandler() IMsgHandle       // 获取当前连接绑定的消息管理器

	GetConnID() string   // GetConnID  //获取当前连接ID
	GetWorkerID() uint32 // Get Worker ID（获取workerId）

	Set(key string, value interface{})  // Set 设置链接属性
	Get(key string) (interface{}, bool) // Get 获取链接属性

	GetLastActivityTime() time.Time // 获取最后活跃时间
	GetMsgTotal() uint64            // 获取该链接处理的消息总数
	GetAccOnStats() bool            // 获取设备点火状态
	SetAccOnStats(bool)             // 设置设备点火状态
	GetAccOnXtTimes() uint64        // 获取设备点火下的心跳次数
	AddAccOnXtTimes(num uint64)     // 增加设备点火下的心跳次数
	SetAccOnXtTimes(num uint64)     // 设置设备点火下的心跳次数
	GetXtTimes() uint64             // 获取设备的心跳次数
	AddXtTimes(num uint64)          // 增加设备的心跳次数
	SetImei(imei string)            // 设置设备imei
	GetImei() string                // 获取设备imei
	UpdateActivity()                // 更新最后活跃时间
	GetLastDeviceStatus() int       // 获取最后设备状态
	SetLastDeviceStatus(int)        // 设置最后设备状态
	GetContext() context.Context    // 获取上下文

}
