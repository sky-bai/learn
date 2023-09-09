package znet

import (
	"backend/bin/tcp/ziface"
	"backend/lib/logger"
	"backend/lib/tools"
	"context"
	"errors"
	"io"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"
)

var _ ziface.IConnection = (*Connection)(nil)

// 这段代码是在Go语言中进行接口断言的语法。它的含义是将`(*Connection)(nil)`转换为`ziface.IConnection`接口类型，并将其赋值给一个匿名的变量`_`。
// 在这里，`(*Connection)(nil)`表示一个`*Connection`类型的空指针。通过将其转换为`ziface.IConnection`接口类型，我们可以判断`*Connection`类型是否实现了`ziface.IConnection`接口。
// 这种用法通常在编译时进行接口的静态类型检查，确保`*Connection`类型实现了`ziface.IConnection`接口的所有方法。如果未实现，则会在编译时产生错误。
// 需要注意的是，由于在代码中使用了匿名变量`_`，该变量不会被使用，仅仅是为了进行接口断言而存在。

type Connection struct {
	//当前Conn属于哪个Server
	TcpServer ziface.IServer //当前conn属于哪个server，在conn初始化的时候添加即可

	//当前连接的socket TCP套接字
	conn *net.TCPConn
	//当前连接的ID
	connId string

	//当前连接的关闭状态
	isClosed bool
	// 当前链接是属于哪个Connection Manager的
	connManager ziface.IConnManager
	//消息管理MsgId和对应处理方法的消息管理模块
	MsgHandler ziface.IMsgHandle

	// The workerid responsible for handling the link
	// 负责处理该链接的workerid
	workerID uint32

	// 当前连接断开时的Hook函数
	onConnStop func(conn ziface.IConnection)

	// 告知该链接已经退出/停止的channel
	ctx    context.Context
	cancel context.CancelFunc

	// 用户收发消息的Lock
	msgLock sync.RWMutex
	//链接属性
	property map[string]interface{}
	//保护链接属性修改的锁
	propertyLock sync.Mutex
	// 心跳检测器
	hc ziface.IHeartbeatChecker
	// 最后一次活动时间
	lastActivityTime time.Time

	// 当前链接处理的消息数量
	msgCount uint64

	// 心跳的消息数量
	XtMsgCount uint64

	// 设备当前的状态
	lastDeviceStatus int

	// 设备是否点火在线
	AccOnStatus bool
	// 设备点火的心跳次数
	AccOnXtTimes uint64
	// 设备Imei号
	Imei string

	// 链接开始的时间
	startTime time.Time
}

// NewConnection 创建连接的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connId string, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server, //将隶属的server传递进来
		conn:       conn,
		connId:     connId,
		isClosed:   false,
		MsgHandler: msgHandler,
		// 当建立与客户端的套接字后，那么就会开启两个Goroutine分别处理读数据业务和写数据业务，读写数据之间的消息通过一个Channel传递。
		property: make(map[string]interface{}), //对链接属性map初始化
		hc:       server.GetHeartBeatChecker(),
		// 给链接初始化心跳检测器    链接里面的属性，没有将心跳作为全局变量  通过new方法去获取对象实例，没有使用全局变量的方法去使用
		msgCount: 0,
	}
	c.connManager = server.GetConnMgr() // (将当前的Connection与Server的ConnManager绑定)
	c.onConnStop = server.GetOnConnStop()
	c.MsgHandler = server.GetMsgHandler()
	c.TcpServer.GetConnMgr().Add(c) //将当前新创建的连接添加到ConnManager中

	return c
}

// StartReader 处理conn读数据的Goroutine 这里是链接里面真正读取数据   处理拆包，然后交给MsgHandler处理
func (c *Connection) StartReader() {

	// 1.关闭当前连接 如果读的时候发生异常，那么就关闭当前连接
	defer c.Stop()

	// 2.recover异常
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			errStr := strings.ReplaceAll(string(buf), "\n", "") // 去除换行
			logger.Error(c.ctx, "StartReader error ", errStr, err)
		}
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:

			// 定义一个对象把封包拆包的逻辑给抽出来，也可以直接写函数
			// 1.创建拆包解包的对象
			dp := NewDataPack()

			newMsg, err := dp.Unpack(c)
			if err != nil {
				// 服务器Accept后，执行ReadFull，此时goroutine会进入阻塞。
				//1, 若client主动调用conn.Close()，则服务器会收到err == EOF。
				//2, 若client主动发送若干字节（少于ReadFull中的buffer），并立即调用conn.Close()，服务器收到EOF Unexpected。
				//3, 若client主动发送恰好等于buffer的字节，并立即调用conn.Close()，服务器收到err == nil, 这种情况下服务器再次调用ReadFull, 会立即收到err == EOF。
				//4, 若client由于直接杀死进程，则服务器收到err == ErrClosed。
				//5, 若client主动关闭链接，就会报connection reset by peer。

				if errors.Is(err, net.ErrClosed) {
					logger.Errorf(c.ctx, "StartReader net.ErrClosed %v", err)
					return
				}

				if errors.Is(err, io.EOF) {
					logger.Errorf(c.ctx, "StartReader io.EOF %v", err)
					return
				}

				// 如果是客户端主动关闭链接，那么就不需要打印日志
				if strings.Contains(err.Error(), "connection reset by peer") {
					logger.Errorf(c.ctx, "StartReader connection reset by peer msg head error %v", err)
					return
				}

				// 如果消息长度为空
				if errors.Is(err, MsgNil) {
					logger.Errorf(c.ctx, "StartReader MsgNil error %v", err)
					return
				}
				// 如果读取失败，那么直接退出该链接
				// 这里要return 才能退出 执行server的stop方法
				logger.Errorf(c.ctx, "StartReader read msg head error %v", err)
				return
			}

			reqNew := Request{
				conn: c,
				msg:  newMsg,
			}

			// 10.将读取到的数据，发送给工作池
			c.MsgHandler.SendMsgToTaskQueue(&reqNew)
		}
	}
}

func (c *Connection) UpdateActivity() {
	c.lastActivityTime = time.Now()
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			errStr := strings.ReplaceAll(string(buf), "\n", "") // 去除换行
			logger.Errorf(c.ctx, "Connection Start ", errStr, err)
		}
	}()

	c.ctx, c.cancel = context.WithCancel(context.Background())
	spanID := tools.RandomStr(10)
	c.ctx = context.WithValue(c.ctx, "spanID", spanID)

	// 占用workerId
	workerID, err := c.MsgHandler.UseWorker(c)
	if err != nil {
		logger.Errorf(c.ctx, "Connection Start UseWorker error %v", err)
		return
	}
	c.workerID = workerID

	// 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.finalizer()
		c.MsgHandler.FreeWorker(c.workerID)
		return
	}
}

func (c *Connection) finalizer() {
	// 0.调用该链接关闭回调函数
	c.callOnConnStop()

	// 1.如果用户注册了该链接的	关闭回调业务，那么在此刻应该显示调用
	c.msgLock.Lock()
	defer c.msgLock.Unlock()

	// 2.如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	// 3.关闭socket链接
	_ = c.conn.Close()

	// 4.关闭链接绑定的心跳检测器
	if c.hc != nil {
		c.hc.Stop(c)
	}

	// 5.将链接从连接管理器中删除
	if c.connManager != nil {
		c.connManager.Remove(c)
	}

	// 6.设置标志位
	c.isClosed = true

}

// Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) callOnConnStop() {
	if c.onConnStop != nil {
		c.onConnStop(c)
	}
}

// Set 设置链接属性
func (c *Connection) Set(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if c.property == nil {
		c.property = make(map[string]interface{})
	}
	c.property[key] = value
}

// Get 获取链接属性
func (c *Connection) Get(key string) (interface{}, bool) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if value, ok := c.property[key]; ok {
		return value, true
	}

	return nil, false
}

func (c *Connection) Write(data string) error {

	// 1.检测链接是否已经关闭
	if c.isClosed == true {
		return errors.New("connection closed when send buff msg")
	}

	// 2.将data封包
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage("", []byte(data)))
	if err != nil {
		logger.Errorf(c.ctx, "Pack error msg data = %d", data)
		return errors.New("Pack error msg ")
	}

	// 3.写回客户端
	if _, err = c.conn.Write(msg); err != nil {
		logger.Errorf(c.ctx, "Send Buff Data error:, %v", err)
		return err
	}
	return nil
}

func (c *Connection) GetLastDeviceStatus() int {
	return c.lastDeviceStatus
}
func (c *Connection) SetLastDeviceStatus(status int) {
	c.lastDeviceStatus = status
}

func (c *Connection) GetLastActivityTime() time.Time {
	return c.lastActivityTime
}

func (c *Connection) SetAccOnStats(accOn bool) {
	c.AccOnStatus = accOn
}

func (c *Connection) AddAccOnXtTimes(i uint64) {
	c.AccOnXtTimes += i
}

func (c *Connection) SetAccOnXtTimes(num uint64) {
	c.AccOnXtTimes = num
}

func (c *Connection) SetImei(imei string) {
	c.Imei = imei
}

func (c *Connection) GetImei() string {
	return c.Imei
}

// GetMsgTotal 获取该链接的消息总数
func (c *Connection) GetMsgTotal() uint64 {
	return c.msgCount
}

// GetAccOnStats 获取设备是否点火状态
func (c *Connection) GetAccOnStats() bool {
	return c.AccOnStatus
}

// GetAccOnXtTimes 获取设备点火状态下的心跳次数
func (c *Connection) GetAccOnXtTimes() uint64 {
	return c.AccOnXtTimes
}

// GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

// GetConnID 获取当前连接ID
func (c *Connection) GetConnID() string {
	return c.connId
}

func (c *Connection) GetWorkerID() uint32 {
	return c.workerID
}

// GetXtTimes 获取设备的心跳次数
func (c *Connection) GetXtTimes() uint64 {
	return c.XtMsgCount
}

// AddXtTimes 增加设备的心跳次数
func (c *Connection) AddXtTimes(num uint64) {
	c.XtMsgCount += num
}

func (c *Connection) GetHeartBeat() ziface.IHeartbeatChecker {
	return c.hc
}

func (c *Connection) GetMsgHandler() ziface.IMsgHandle {
	return c.MsgHandler
}

func (c *Connection) GetContext() context.Context {
	return c.ctx
}
