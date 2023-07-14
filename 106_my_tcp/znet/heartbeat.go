package znet

import (
	"backend/bin/tcp/tcpParser"
	"backend/bin/tcp/ziface"
	"backend/lib/logger"
	"context"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

var _ ziface.IHeartbeatChecker = (*HeartbeatChecker)(nil)

// 设备状态
const (
	slots = 420 // 设置任务的最大延迟时间为7分钟 时间轮的槽数需要根据任务的延迟时间来确定

	suoShi         = 3
	suoShiInterval = 3*time.Minute + 2*time.Second

	DevStatusAccOnOnline = 1                                // 点火在线(TCP)
	DevAccOnInterval     = 10 * time.Second                 // 设备点火在线状态上报频率
	CheckAccOnInterval   = DevAccOnInterval + 2*time.Minute // 检测设备心跳频率 点火在线状态上报频率 + 2s容错时长

	DevStatusAccOffOnline     = 2                                       // 熄火在线(TCP)
	DevAccOffOnlineInterval   = 5 * time.Minute                         // 设备熄火在线状态上报频率
	CheckAccOffOnlineInterval = DevAccOffOnlineInterval + 2*time.Minute // 熄火在线状态上报频率 + 2s容错时长

	DevStatusShutdown           = 3                             // 关机
	DevStatusAccOffOffline      = 4                             // 熄火离线
	DevStatusOffline            = 5                             //离线（异常）
	DevStatusLowPowerShutdown   = 6                             // 低电关机
	DevStatusParkToSecurity     = 7                             // 缩时录影(TCP)
	CheckParkToSecurityInterval = 3*time.Minute + 2*time.Minute // 缩时录影状态上报频率 + 2s容错时长
	DevStatusTimeoutShutdown    = 8                             // 超时关机
)

// HeartbeatChecker 心跳检测器 通过时间轮算法去检测心跳

// 心跳检测器会设备上报的心跳时间间隔基础上+2秒，去判断检测时间和设备最后一次活跃时间是否大于时间间隔，如果大于，就认为设备不在线

type HeartbeatChecker struct {
	interval         time.Duration           // 心跳检测时间间隔
	onRemoteNotAlive ziface.OnRemoteNotAlive //用户自定义的远程连接不存活时的处理方法
	timingWheel      *collection.TimingWheel // 时间轮
	connManager      ziface.IConnManager     // 连接管理器
}

// NewHeartbeatChecker 创建心跳检测器
func NewHeartbeatChecker(interval time.Duration) (ziface.IHeartbeatChecker, error) {
	heartbeat := &HeartbeatChecker{
		interval: interval,
	}

	timingWheel, err := collection.NewTimingWheel(interval, slots, func(k, v any) {
		// 1. 获取连接
		conn, ok := k.(ziface.IConnection)
		if !ok { // 如果不是就直接返回了
			logger.Errorf(context.Background(), "not uint64 type")
			return
		}

		// 2. 如果设备在线，判断最新的心跳时间是否超过了设定的时间间隔
		// 当前时间减去上次链接的活跃时间，如果大于设备上传时间间隔，就认为设备不在线

		// 3.获取链接的最新的心跳时间 对该时间进行判断
		checkInterval := GetHeartbeatIntervalByStatus(conn.GetLastDeviceStatus())

		// 4. 如果设备在线，判断最新的心跳时间是否超过了设定的时间间隔
		if conn.GetLastActivityTime().Add(checkInterval).Before(time.Now()) {
			// 3.1.1 如果超过了，就将该设备设置为离线
			// 远程连接不存活时的处理方法
			if heartbeat.onRemoteNotAlive != nil {
				heartbeat.onRemoteNotAlive(conn)
			}
			conn.Stop()
		} else {
			// 3.1.2 如果没有超过,就更新定时任务
			conn.GetHeartBeat().ResetConnExpire(conn, conn.GetLastActivityTime().Add(checkInterval).Sub(time.Now()))
		}

		//if heartbeat.onRemoteNotAlive != nil {
		//	// TODO 这儿为什么又取了一次？
		//	// 存入的string connId 相当于我可以存conn，然后直接调用
		//	conn, err = heartbeat.connManager.Get(connId)
		//	if err != nil {
		//		logger.Errorf(context.Background(), "connManager get conn error: %s", err.Error())
		//		return
		//	}
		//	// TODO 这儿不应该是判断链接最近一次心跳的时间，然后判断是否需要断开链接吗？
		//	// 我看go-zero的时间轮在创建的时候就写好到时间点的实际业务逻辑，然后任务随时加入，或者是移除，到时间点就会执行 移动任务和新加任务都是往channel里面发 这样的话，我每次只是移动，不用判断，到了时间点就直接执行回调操作。
		//
		//	// TODO 心跳指令只需要维护最近在线时间即可，不需要管时间轮
		//	// 1.创建任务，2.移动任务，3.到时间点执行任务
		//	// 如果这样做的话好像不需要最近在线时间了，之前好像是要判断是否在当前设备的状态时间之内，如果超出就断开链接，那么每一次时间点执行的任务就会很多。
		//
		//	logger.Infof(context.Background(), "connect not alive imei:%s, connId: %d, lastActivityTime: %s, now: %s", conn.GetImei(), connId, conn.GetLastActivityTime().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
		//	heartbeat.onRemoteNotAlive(conn)
		//}

	})

	if err != nil {
		// TODO panic后，当前goroutine会立即停止，后面的return应该执行不到了
		// 是的
		panic(err)
	}

	heartbeat.timingWheel = timingWheel

	return heartbeat, nil
}

// BindConnManger 绑定链接Manager
func (h *HeartbeatChecker) BindConnManger(connManager ziface.IConnManager) {
	h.connManager = connManager
}

// Set 当前链接启动心跳检测 TODO 没有看见这个方法在哪儿被调用呢

// 之前是在链接创建的时候set 然后xt的时候move，我好像删掉了，应该要加回去
func (h *HeartbeatChecker) Set(conn ziface.IConnection, interval, afterTime time.Duration) {
	err := h.timingWheel.SetTimer(conn, interval, afterTime)
	if err != nil {
		logger.Errorf(context.Background(), " HeartbeatChecker Set: %s", err.Error())
		return
	}
}

// Stop 当前链接停止心跳检测
func (h *HeartbeatChecker) Stop(conn ziface.IConnection) {
	err := h.timingWheel.RemoveTimer(conn)
	if err != nil {
		logger.Errorf(context.Background(), " HeartbeatChecker Stop: %s", err.Error())
		return
	}
}

// ResetConnExpire  重置当前链接的心跳检测 将当前链接的心跳检测时间更新为interval之后 setTimer方法将会被取消
func (h *HeartbeatChecker) ResetConnExpire(conn ziface.IConnection, interval time.Duration) {
	err := h.timingWheel.MoveTimer(conn, interval) // moveTimer 是在当前时间的基础上加上interval 不是之前设置的时间加上interval
	if err != nil {
		logger.Errorf(context.Background(), " HeartbeatChecker ResetConnExpire: %s", err.Error())
		return
	}
}

func (h *HeartbeatChecker) SetOnRemoteNotAlive(f ziface.OnRemoteNotAlive) {
	if f != nil {
		h.onRemoteNotAlive = f
	}
}

// Clone 克隆到一个指定的链接上
func (h *HeartbeatChecker) Clone() ziface.IHeartbeatChecker {

	heartbeat := &HeartbeatChecker{
		interval:         h.interval,
		onRemoteNotAlive: h.onRemoteNotAlive,
		connManager:      nil, //绑定的链接需要重新赋值
	}

	return heartbeat
}

// GetHeartbeatInterval 根据status获取心跳检测的时间间隔
func GetHeartbeatInterval(msg []byte) (int, time.Duration) {

	status := tcpParser.ParseXt(string(msg)).Status

	// 设备状态的变化导致设备上传的心跳频率发生变化，需要更新心跳检测的时间间隔
	var checkInterval time.Duration
	switch status {
	case DevStatusAccOnOnline: // 设备在线
		checkInterval = CheckAccOnInterval
	case DevStatusAccOffOnline: // 熄火在线
		checkInterval = CheckAccOffOnlineInterval
	case DevStatusParkToSecurity: // 缩时录影
		checkInterval = CheckParkToSecurityInterval
	}

	return status, checkInterval
}

// GetHeartbeatIntervalByStatus 根据status获取心跳检测的时间间隔
func GetHeartbeatIntervalByStatus(deviceStatus int) time.Duration {

	// 设备状态的变化导致设备上传的心跳频率发生变化，需要更新心跳检测的时间间隔
	var checkInterval time.Duration
	switch deviceStatus {
	case DevStatusAccOnOnline: // 设备在线
		checkInterval = CheckAccOnInterval
	case DevStatusAccOffOnline: // 熄火在线
		checkInterval = CheckAccOffOnlineInterval
	case DevStatusParkToSecurity: // 缩时录影
		checkInterval = CheckParkToSecurityInterval
	}

	return checkInterval
}
