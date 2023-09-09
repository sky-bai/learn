package zmetrics

const (
	GaugeConnectionTotalName string = "connection_online_total"
	GaugeConnectionTotalHelp string = "All Online Connections Group By (Address, Name) ( 不同Server的链接总数,根据(Address, Name) 分组)"

	GaugeTaskTotalName string = "task_total"
	GaugeTaskTotalHelp string = "All Task Total Group By (Address, Name, WorkerID) ( 已经处理的数据任务总数,根据(Address, Name, WorkerID)分组)"

	GaugeRouterScheduleTotalName string = "router_schedule_total"
	GaugeRouterScheduleTotalHelp string = "Router Schedule Total Group By (Address, Name, WorkerID, MsgID) ( 路由调度的Handler总数,根据(Address, Name, WorkerID, MsgID)分组)"

	HistogramRouterScheduleDurationName string = "router_schedule_duration"
	HistogramRouterScheduleDurationHelp string = "Router Schedule Duration Group By (Address, Name, WorkerID, MsgID) ( 路由调度的Handler耗时,根据(Address, Name, WorkerID, MsgID)分组)"

	HistogramConnDurationName string = "connect_duration"
	HistogramConnDurationHelp string = "Connect Duration Group By (Address, Name) ( 链接耗时,根据(Address, Name)分组)"
)

const (
	LabelAddress  string = "address"
	LabelName     string = "name"
	LabelWorkerID string = "worker_id"
	LabelMsgID    string = "msg_id"
)
