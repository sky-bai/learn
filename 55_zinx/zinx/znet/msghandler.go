package znet

import (
	"fmt"
	"learn/55_zinx/zinx/utils"
	"learn/55_zinx/zinx/ziface"
	"strconv"
)

// 管理消息类型和路由的模块

type MsgHandle struct {
	// 消息类型以数字去保存 没有以string去保存
	// todo 用map去存对应关系，还能用什么数据结构去存储昵？
	Apis           map[uint32]ziface.IRouter //存放每个MsgId 所对应的处理方法的map属性
	WorkerPoolSize uint32                    //业务工作Worker池的数量
	TaskQueue      []chan ziface.IRequest    //Worker负责取任务的消息队列 // 管理多个channel request 多批次请求 开多个通道 然后开多个worker去处理 构建多个消息队列 创建出来的key就是worker的id，也就是消息队列的编号
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		//一个worker对应一个queue
		TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	//执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	//2 添加msg与api的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}

// StartOneWorker 启动一个Worker工作流程  具体worker的工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select { // 不断处理链接中的数据 这里是用channel把数据保存起来的
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// StartWorkerPool 启动worker工作池 这里是我到底要开好多个worker
func (mh *MsgHandle) StartWorkerPool() {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	mh.TaskQueue[workerID] <- request // 把消息放到对应的消息队列中
}

// StartWorkerPool()方法是启动Worker工作池，这里根据用户配置好的WorkerPoolSize的数量来启动，然后分别给每个Worker分配一个TaskQueue，然后用一个goroutine来承载一个Worker的工作业务。
//
// StartOneWorker()方法就是一个Worker的工作业务，每个worker是不会退出的(目前没有设定worker的停止工作机制)，会永久的从对应的TaskQueue中等待消息，并处理。

// 接下来我们就需要给Zinx添加消息队列和多任务Worker机制了。
// 我们可以通过worker的数量来限定处理业务的固定goroutine数量，而不是无限制的开辟Goroutine，
// 虽然我们知道go的调度算法已经做的很极致了，但是大数量的Goroutine依然会带来一些不必要的环境切换成本，
// 这些本应该是服务器应该节省掉的成本。我们可以用消息队列来缓冲worker工作的数据。
