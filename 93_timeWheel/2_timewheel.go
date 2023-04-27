package main

import (
	"fmt"
	"sync"
	"time"
)

// 层级时间轮是一种高效的定时器实现方式，可以满足复杂的时间调度需求。下面是用Go语言实现层级时间轮的示例代码：

// TimeWheelSlot 时间槽
type TimeWheelSlot struct {
	Jobs []*Job
}

// TimeWheel 时间轮
type TimeWheel struct {
	Slots      []*TimeWheelSlot // 时间轮槽 有多大
	WheelSize  int              // 时间轮的大小
	Interval   time.Duration    // 时间轮转动一次的时间间隔
	Tick       int              // 时间轮当前的刻度
	keysetLock *sync.RWMutex    // 锁
}

// 如何检测是否有效，或者是有问题

// Job 定时器任务
type Job struct {
	ID      int               // 定时任务ID
	Time    time.Time         // 参照时间
	Target  interface{}       // 定时任务携带的数据
	Exptime time.Duration     // 经过多长时间后触发定时任务
	Cigo    *time.Timer       // 定时器引用
	Handler func(interface{}) // 处理函数
}

// NewSlot 创建一个时间槽
func NewSlot() *TimeWheelSlot {
	return &TimeWheelSlot{}
}

// NewTimeWheel 创建一个时间轮
func NewTimeWheel(interval time.Duration, wheelSize int) *TimeWheel {
	return &TimeWheel{
		Slots:      make([]*TimeWheelSlot, wheelSize),
		Interval:   interval,
		Tick:       0,
		WheelSize:  wheelSize,
		keysetLock: new(sync.RWMutex),
	}
}

// AddJob 增加一个定时任务
func (tw *TimeWheel) AddJob(job *Job) {
	if job.Exptime < tw.Interval {
		go job.Handler(job.Target)
		return
	}

	pos, exptime := tw.getPositionAndExptime(job)
	job.Exptime = exptime
	tw.Slots[pos].Jobs = append(tw.Slots[pos].Jobs, job)
}

// getPositionAndExptime 计算定时任务要插入到轮中哪个位置
func (tw *TimeWheel) getPositionAndExptime(job *Job) (pos int, exptime time.Duration) {
	exptime = job.Exptime

	if exptime < tw.Interval*time.Duration(tw.WheelSize) {
		pos = (tw.Tick + int(exptime/tw.Interval)) % tw.WheelSize
	} else {
		pos = (tw.Tick + tw.WheelSize - 1) % tw.WheelSize
	}

	return pos, exptime
}

// RemoveJob 删除一个已经存在的定时任务
func (tw *TimeWheel) RemoveJob(job *Job) {
	pos, _ := tw.getPositionAndExptime(job)

	for index, j := range tw.Slots[pos].Jobs {
		if j.ID == job.ID {
			tw.Slots[pos].Jobs = append(tw.Slots[pos].Jobs[:index], tw.Slots[pos].Jobs[index+1:]...)
			return
		}
	}
}

// Schedule 调度
func (tw *TimeWheel) Schedule() {
	tw.keysetLock.Lock()
	defer tw.keysetLock.Unlock()

	pos := tw.Tick % tw.WheelSize
	slot := tw.Slots[pos]

	for _, job := range slot.Jobs {
		go job.Handler(job.Target)
	}

	tw.Slots[pos].Jobs = make([]*Job, 0)
	tw.Tick++

	if tw.Tick == tw.WheelSize {
		tw.Tick = 0
	}
}

// Run 运行时间轮定时器
func (tw *TimeWheel) Run() {
	ticker := time.NewTicker(tw.Interval)

	for {
		<-ticker.C
		tw.Schedule()
	}
}

//上述代码中，TimeWheel代表时间轮，它包含一个轮槽数组Slots，表示整个时间轮。Interval表示时间轮转动一次的时间间隔，Tick表示时间轮当前的刻度，WheelSize代表时间轮的大小。keysetLock是一个读写锁，用于保护时间槽的并发访问。
//
//TimeWheelSlot表示一个时间槽，它包含一个Job数组，表示处于该时间槽的所有定时任务。
//
//Job代表一个定时任务，其属性包括ID、Time、Target、Exptime、Cigo和Handler。ID是定时任务的id，Time是任务的参考时间，Target是任务携带的数据，Exptime是经过多长时间后触发定时任务，Cigo是定时器引用，Handler是处理函数。
//
//函数AddJob向时间轮中增加一个定时任务，getPositionAndExptime计算将该任务插入到轮中哪个位置，RemoveJob删除一个已经存在的定时任务，Schedule用于遍历执行时间轮中的任务，而Run则是用于运行整个时间轮定时器。
//
//为了使用这些程序，在main函数中创建一个时间轮并运行：

func main() {
	// 创建时间轮
	tw := NewTimeWheel(time.Millisecond*100, 360)

	// 启动时间轮
	go tw.Run()

	// 向时间轮中添加任务
	tw.AddJob(&Job{
		ID:      1,
		Time:    time.Now(),
		Target:  "test",
		Exptime: time.Millisecond * 200,
		Handler: func(target interface{}) {
			fmt.Printf("任务执行了, 参数:%#v\n", target)
		},
	})

	// 暂停一秒钟，等待任务执行
	time.Sleep(time.Second)
}

//这个示例创建一个时间轮，并向其中添加一个任务。添加任务后，程序暂停1秒钟，等待任务执行完毕。
