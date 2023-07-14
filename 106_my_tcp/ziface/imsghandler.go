package ziface

/*
	消息管理抽象层 消息的具体处理方式和路由
*/

type IMsgHandle interface {
	AddHandler(msgId string, handler func(request IRequest)) //为消息添加具体的处理逻辑
	StartWorkerPool()                                        //启动worker工作池
	SendMsgToTaskQueue(request IRequest)                     //将消息交给TaskQueue,由worker进行处理

	UseWorker(IConnection) (uint32, error) // Use worker ID
	FreeWorker(workerID uint32)            // Free worker ID

	GetApisHandler() map[string]func(request IRequest) // Get all apis handler
}
