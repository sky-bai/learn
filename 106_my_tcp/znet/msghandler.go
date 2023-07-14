package znet

import (
	"backend/bin/tcp/ziface"
	"backend/config"
	"backend/lib/logger"
	"context"
	"errors"
	"strings"
	"sync"
)

var _ ziface.IMsgHandle = (*MsgHandle)(nil)

var ErrNoWorker = errors.New("no worker")

type MsgHandle struct {
	Apis           map[string]func(request ziface.IRequest) // 存放每个MsgId 所对应的处理方法的map属性
	WorkerPoolSize uint32                                   // 业务工作Worker池的数量	TODO 是不是会引入频繁切换worker的问题
	// worker是一个ID， 后面的chan和goroutine创建的时候是初始化好了的，请求来了就分给对应的chan，这里的频繁切换指的是什么？
	TaskQueue []chan ziface.IRequest // Worker负责取任务的消息队列

	freeWorkers  map[uint32]struct{} // 空闲worker集合，用于平衡消息分发
	freeWorkerMu sync.Mutex
}

func NewMsgHandle() *MsgHandle {

	handle := &MsgHandle{
		Apis: make(map[string]func(request ziface.IRequest)),
		// 这个ID对应的中间件
		WorkerPoolSize: config.Config.Tcp.WorkerPoolSize,
		// WorkerPoolSize:作为工作池的数量，因为TaskQueue中的每个队列应该是和一个Worker对应的，所以我们在创建TaskQueue中队列数量要和Worker的数量一致.

		//一个worker对应一个queue
		TaskQueue: make([]chan ziface.IRequest, config.Config.Tcp.WorkerPoolSize),
		// TaskQueue真是一个Request请求信息的channel集合。用来缓冲提供worker调用的Request请求信息，worker会从对应的队列中获取客户端的请求数据并且处理掉。

		freeWorkerMu: sync.Mutex{},
	}

	// 预先创建好Worker，每个Worker用一个go来承载
	handle.freeWorkers = make(map[uint32]struct{}, config.Config.Tcp.WorkerPoolSize)
	for i := uint32(0); i < config.Config.Tcp.WorkerPoolSize; i++ {
		handle.freeWorkers[i] = struct{}{}
	}

	return handle
}

// AddHandler 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddHandler(msgId string, handler func(request ziface.IRequest)) {
	//1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + msgId)
	}

	//2 添加msg与api的绑定关系
	mh.Apis[msgId] = handler
}

// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 1.读取消息的第三位，是该消息的类型 TODO 需要考虑数组越界的异常
	// 好的
	arrRequest := strings.Split(string(request.GetData()), ",")
	if len(arrRequest) < 3 {
		logger.Warnf(context.Background(), "SendMsgToTaskQueue arrRequest len < 3, arrRequest:%s", arrRequest)
		return
	}
	requestMsgType := arrRequest[2]

	// 2.如果该消息类型没有在后台注册过,则不处理
	apis := mh.GetApisHandler()
	if _, ok := apis[requestMsgType]; !ok {
		logger.Warnf(context.Background(), "This message type has not been registered in the background, requestMsgType:%s,string(req.GetData()):%s", requestMsgType, request.GetData())
		return
	}

	// 设置消息类型
	request.SetMsgType(requestMsgType)

	// 得到需要处理此条连接的workerID
	workerId := request.GetConnection().GetWorkerID()

	// 拿到可用的workid，然后去找对应的队列，最后将消息发送给队列
	// 将请求消息发送给任务队列
	mh.TaskQueue[workerId] <- request
}

// StartWorkerPool 启动worker工作池
func (mh *MsgHandle) StartWorkerPool() {
	//遍历需要启动worker的数量，依此启动
	for i := 0; uint32(i) < mh.WorkerPoolSize; i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, config.Config.Tcp.MaxWorkerTaskLen) // 8是一个worker对应的消息队列的长度
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(mh.TaskQueue[i]) // 有多少worker就开多少个goroutine
	}
}

// StartOneWorker 启动一个Worker工作流程 一个worker对应一个taskQueue
func (mh *MsgHandle) StartOneWorker(taskQueue chan ziface.IRequest) {
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:

			mh.doMsgHandler(request)
		}
	}
}

// doMsgHandler 立即以非阻塞方式处理消息
func (mh *MsgHandle) doMsgHandler(request ziface.IRequest) {

	handler, ok := mh.Apis[request.GetMsgType()]
	if !ok {
		logger.Errorf(context.Background(), "api msgId = %d is not FOUND!", request.GetMsgType())
		return
	}

	handler(request)
}

// UseWorker 占用workerID
func (mh *MsgHandle) UseWorker(conn ziface.IConnection) (uint32, error) {
	// 改成空闲hashmap，以进行绝对的负载均衡，仅适用于新算法，因为新算法不会有重叠
	// 新算法应该有两个封装函数，务必确保释放和回收的数量一致
	mh.freeWorkerMu.Lock()
	defer mh.freeWorkerMu.Unlock()

	// 这个map存的是可用的workId
	// 从空闲worker中获取一个workerId，之前是50w的map，遍历出一个就减少一个，然后直接返回
	// 获取map中的一个key
	for k := range mh.freeWorkers {
		delete(mh.freeWorkers, k)
		return k, nil
	}

	// TODO 这儿是不是不对，执行到这儿表明freeWorkers已经使用完了，应该跑错才对，为什么是随机抛一个出去呢？
	// 之前在创建链接的时候只是打印该扩容了，然后将超过的链接还是处理了的，之前想的是超出的链接按照链接ID去轮询去分发的，链接ID是雪花就可以获取int类型和string类型，然后改成了随机字符串，获取不到int，就没有求余轮询了，
	return 0, ErrNoWorker
}

// FreeWorker 释放workerId
func (mh *MsgHandle) FreeWorker(workerID uint32) {
	mh.freeWorkerMu.Lock()
	defer mh.freeWorkerMu.Unlock()

	mh.freeWorkers[workerID] = struct{}{}

}

func (mh *MsgHandle) GetApisHandler() map[string]func(request ziface.IRequest) {
	return mh.Apis
}
