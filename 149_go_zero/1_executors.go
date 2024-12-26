package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/lang"
	"sync"
	"time"
)

// 批处理任务

type TaskContainer interface {
	AddTask(task any) bool
	Execute(tasks any)
	RemoveAll() any
}

// 实现类

type bulkContainer struct {
	tasks    []any
	maxTasks int
}

func (b *bulkContainer) AddTask(task any) bool {
	b.tasks = append(b.tasks, task)
	return len(b.tasks) >= b.maxTasks
}

func (b *bulkContainer) RemoveAll() any {
	tasks := b.tasks
	b.tasks = nil
	return tasks
}

func (b *bulkContainer) Execute(tasks any) {
	vals := tasks.([]any)
	b.execute(vals)
}

func (b *bulkContainer) execute(tasks []any) {
	fmt.Println("bulkContainer execute")
}

type chunkContainer struct {
	tasks        []any
	execute      Execute
	size         int
	maxChunkSize int
}

func (c *chunkContainer) AddTask(task any) bool {
	c.tasks = append(c.tasks, task)
	return len(c.tasks) >= c.maxChunkSize
}

type chunk struct {
	val  any
	size int
}

func (c *chunkContainer) RemoveAll() any {
	tasks := c.tasks
	c.tasks = nil
	c.size = 0
	return tasks
}

type PeriodicalExecutor struct {
	commander   chan any
	interval    time.Duration
	container   TaskContainer
	waitGroup   sync.WaitGroup
	wgBarrier   Barrier
	confirmChan chan lang.PlaceholderType
	inflight    int32
	guarded     bool
	newTicker   func(duration time.Duration) *time.Ticker
	lock        sync.Mutex

	// 先看使用的方法
}

// 资源同步访问的屏障机制
// 需要对资源对访问是线程安全的

// 实现就是锁

type Barrier struct {
	lock sync.Mutex
}

// Guard guards the given fn on the resource. 包装对资源的安全访问
func (b *Barrier) Guard(fn func()) {
	Guard(&b.lock, fn)
}

// Guard guards the given fn with lock.
func Guard(lock sync.Locker, fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}

func (pe *PeriodicalExecutor) Add(task any) {
	pe.lock.Lock()
	defer func() {
		// 一开始为 false
		if !pe.guarded {
			pe.guarded = true
			// 是否只需要关注使用就可以了
		}
	}()
}

// 去重处理 singleFlight
//
