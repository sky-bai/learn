package main

// Job 实际执行任务的类型
type Job struct {
	Payload int
}

// JobQueue 任务队列
var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func main() {

}
