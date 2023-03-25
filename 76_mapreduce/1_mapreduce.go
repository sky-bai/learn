package mapreduce

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

const (
	defaultWorkers = 16
	minWorkers     = 1
)

var (
	// ErrCancelWithNil is an error that mapreduce was cancelled with nil.
	ErrCancelWithNil = errors.New("mapreduce cancelled with nil")
	// ErrReduceNoOutput is an error that reduce did not output a value.
	ErrReduceNoOutput = errors.New("reduce not writing value")
)

type (
	// ForEachFunc is used to do element processing, but no output. mapper 是对每个元素进行处理
	ForEachFunc[T any] func(item T) // 要处理的每个元素
	// GenerateFunc is used to let callers send elements into source.
	GenerateFunc[T any] func(source chan<- T) // 传入一个只写的chan 类型为任意类型 因为是流数据所有用管道 接受所有元素的通道
	// MapFunc is used to do element processing and write the output to writer.
	MapFunc[T, U any] func(item T, writer Writer[U]) // 某个元素，一个实现了Writer接口的任意类型
	// MapperFunc is used to do element processing and write the output to writer, 处理每个元素 并写入到writer 单独处理每个元素
	// use cancel func to cancel the processing.
	MapperFunc[T, U any] func(item T, writer Writer[U], cancel func(error))
	// ReducerFunc is used to reduce all the mapping output and write to writer,
	// use cancel func to cancel the processing.
	ReducerFunc[U, V any] func(pipe <-chan U, writer Writer[V], cancel func(error)) // 聚合所有的map的结果，写入到writer中
	// VoidReducerFunc is used to reduce all the mapping output, but no output.
	// Use cancel func to cancel the processing.
	VoidReducerFunc[U any] func(pipe <-chan U, cancel func(error))
	// Option defines the method to customize the mapreduce.
	Option func(opts *mapReduceOptions)

	mapperContext[T, U any] struct {
		ctx       context.Context
		mapper    MapFunc[T, U]
		source    <-chan T
		panicChan *onceChan
		collector chan<- U
		doneChan  <-chan struct{}
		workers   int
	}

	mapReduceOptions struct {
		ctx     context.Context
		workers int
	}

	// Writer interface wraps Write method.
	Writer[T any] interface {
		Write(v T)
	}
)

// Finish runs fns parallelly, cancelled on any error.
func Finish(fns ...func() error) error {
	if len(fns) == 0 {
		return nil
	}

	return MapReduceVoid(func(source chan<- func() error) {
		for _, fn := range fns {
			source <- fn
		}
	}, func(fn func() error, writer Writer[any], cancel func(error)) {
		if err := fn(); err != nil {
			cancel(err)
		}
	}, func(pipe <-chan any, cancel func(error)) {
	}, WithWorkers(len(fns)))
}

// FinishVoid runs fns parallelly.
func FinishVoid(fns ...func()) {
	if len(fns) == 0 {
		return
	}

	ForEach(func(source chan<- func()) {
		for _, fn := range fns {
			source <- fn
		}
	}, func(fn func()) {
		fn()
	}, WithWorkers(len(fns)))
}

// ForEach maps all elements from given generate but no output.
func ForEach[T any](generate GenerateFunc[T], mapper ForEachFunc[T], opts ...Option) {
	options := buildOptions(opts...)
	panicChan := &onceChan{channel: make(chan any)}
	source := buildSource(generate, panicChan)
	collector := make(chan any, options.workers)
	done := make(chan struct{})

	go executeMappers(mapperContext[T, any]{
		ctx: options.ctx,
		mapper: func(item T, writer Writer[any]) {
			mapper(item)
		},
		source:    source,
		panicChan: panicChan,
		collector: collector,
		doneChan:  done,
		workers:   options.workers,
	})
	// 其实就是一个事件在做的时候，如果其他优先级更高的事情发生了就不执行该事情了。
	for {
		select {
		// 两个事件
		case v := <-panicChan.channel:
			panic(v)
		case _, ok := <-collector:
			if !ok {
				return
			}
		}
	}
}

// MapReduce maps all elements generated from given generate func,
// and reduces the output elements with given reducer.
func MapReduce[T, U, V any](generate GenerateFunc[T], mapper MapperFunc[T, U], reducer ReducerFunc[U, V], opts ...Option) (V, error) {
	panicChan := &onceChan{channel: make(chan any)} // 创造一个panic chan
	source := buildSource(generate, panicChan)
	return mapReduceWithPanicChan(source, panicChan, mapper, reducer, opts...)
}

// MapReduceChan maps all elements from source, and reduce the output elements with given reducer.
func MapReduceChan[T, U, V any](source <-chan T, mapper MapperFunc[T, U], reducer ReducerFunc[U, V],
	opts ...Option) (V, error) {
	panicChan := &onceChan{channel: make(chan any)}
	return mapReduceWithPanicChan(source, panicChan, mapper, reducer, opts...)
}

// MapReduceChan maps all elements from source, and reduce the output elements with given reducer.
func mapReduceWithPanicChan[T, U, V any](source <-chan T, panicChan *onceChan, mapper MapperFunc[T, U], reducer ReducerFunc[U, V], opts ...Option) (val V, err error) {
	// 1.创建option 作用是什么 ？ 用于控制mapreduce的行为
	options := buildOptions(opts...)
	// output is used to write the final result
	// 2.创建最终的结果chan 作用是用于写入最终的结果
	output := make(chan V) // 最终的结果
	// 3.创建panic chan
	defer func() {
		// reducer can only write once, if more, panic // 聚合函数只能写一次，如果多次写入，会panic
		for range output {
			panic("more than one element written in reducer")
		}
	}()

	// 4.创建collector chan 作用是用于写入map的结果
	// collector is used to collect data from mapper, and consume in reducer
	collector := make(chan U, options.workers)

	// 5.创建done chan 作用是用于通知mapper和reducer退出
	// if done is closed, all mappers and reducer should stop processing
	done := make(chan struct{}) // 创建一个完成的信号
	// 6.创建mapper的上下文
	writer := newGuardedWriter(options.ctx, output, done)

	// 7.创建只关闭一次的once
	var closeOnce sync.Once
	// 8.使用原子操作避免数据竞争
	// use atomic.Value to avoid data race
	var retErr atomic.Value
	// 9.创建finish 关闭done output chan 作用是避免内存泄漏
	finish := func() {
		closeOnce.Do(func() {
			close(done)
			close(output)
		})
	}
	// 10.看球不懂 创建一个cancel的方法 作用是 如果有错误了，就关闭done chan 和 output chan 和排干里面的元素
	cancel := once(func(err error) {
		if err != nil {
			retErr.Store(err)
		} else {
			retErr.Store(ErrCancelWithNil)
		}
		// 释放资源
		// 1.排干chan里面的元素
		drain(source)
		// 2.关闭done chan
		finish()
	})

	// 11.异步执行reducer
	go func() {
		defer func() {
			drain(collector)
			if r := recover(); r != nil {
				panicChan.write(r)
			}
			finish()
		}()

		reducer(collector, writer, cancel)
	}()

	go executeMappers(mapperContext[T, U]{
		ctx: options.ctx,
		mapper: func(item T, w Writer[U]) {
			mapper(item, w, cancel)
		},
		source:    source,
		panicChan: panicChan,
		collector: collector,
		doneChan:  done,
		workers:   options.workers,
	})

	select {
	case <-options.ctx.Done():
		cancel(context.DeadlineExceeded)
		err = context.DeadlineExceeded
	case v := <-panicChan.channel:
		panic(v)
	case v, ok := <-output:
		if e := retErr.Load(); e != nil {
			err = e.(error)
		} else if ok {
			val = v
		} else {
			err = ErrReduceNoOutput
		}
	}

	return
}

// MapReduceVoid maps all elements generated from given generate,
// and reduce the output elements with given reducer.
func MapReduceVoid[T, U any](generate GenerateFunc[T], mapper MapperFunc[T, U],
	reducer VoidReducerFunc[U], opts ...Option) error {
	_, err := MapReduce(generate, mapper, func(input <-chan U, writer Writer[any], cancel func(error)) {
		reducer(input, cancel)
	}, opts...)
	if errors.Is(err, ErrReduceNoOutput) {
		return nil
	}

	return err
}

// WithContext customizes a mapreduce processing accepts a given ctx.
func WithContext(ctx context.Context) Option {
	return func(opts *mapReduceOptions) {
		opts.ctx = ctx
	}
}

// WithWorkers customizes a mapreduce processing with given workers.
func WithWorkers(workers int) Option {
	return func(opts *mapReduceOptions) {
		if workers < minWorkers {
			opts.workers = minWorkers
		} else {
			opts.workers = workers
		}
	}
}

func buildOptions(opts ...Option) *mapReduceOptions {
	options := newOptions()
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func buildSource[T any](generate GenerateFunc[T], panicChan *onceChan) chan T {
	source := make(chan T) // 创造一个meta data  chan
	go func() {
		defer func() {
			if r := recover(); r != nil {
				panicChan.write(r)
			}
			close(source)
		}()

		generate(source)
	}()

	return source
}

// drain drains the channel.  // 排干channel
func drain[T any](channel <-chan T) {
	// drain the channel
	for range channel {
	}
}

func executeMappers[T, U any](mCtx mapperContext[T, U]) {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		close(mCtx.collector)
		drain(mCtx.source)
	}()

	var failed int32
	pool := make(chan struct{}, mCtx.workers)
	writer := newGuardedWriter(mCtx.ctx, mCtx.collector, mCtx.doneChan)
	for atomic.LoadInt32(&failed) == 0 {
		select {
		case <-mCtx.ctx.Done():
			return
		case <-mCtx.doneChan:
			return
		case pool <- struct{}{}:
			item, ok := <-mCtx.source
			if !ok {
				<-pool
				return
			}

			wg.Add(1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						atomic.AddInt32(&failed, 1)
						mCtx.panicChan.write(r)
					}
					wg.Done()
					<-pool
				}()

				mCtx.mapper(item, writer)
			}()
		}
	}
}

func newOptions() *mapReduceOptions {
	return &mapReduceOptions{
		ctx:     context.Background(),
		workers: defaultWorkers,
	}
}

func once(fn func(error)) func(error) { // sync.Once 和 sync.Once.Do 的区别 sync.Once 的创建和使用
	once := new(sync.Once)   // 单飞库的使用
	return func(err error) { // 高阶函数创建对象，低阶函数调用对象方法
		once.Do(func() {
			fn(err)
		})
	}
}

type guardedWriter[T any] struct {
	ctx     context.Context
	channel chan<- T
	done    <-chan struct{}
}

func newGuardedWriter[T any](ctx context.Context, channel chan<- T, done <-chan struct{}) guardedWriter[T] {
	return guardedWriter[T]{
		ctx:     ctx,
		channel: channel,
		done:    done,
	}
}

func (gw guardedWriter[T]) Write(v T) {
	select {
	case <-gw.ctx.Done():
		return
	case <-gw.done:
		return
	default:
		gw.channel <- v
	}
}

type onceChan struct {
	channel chan any
	wrote   int32
}

func (oc *onceChan) write(val any) {
	if atomic.AddInt32(&oc.wrote, 1) > 1 {
		return
	}

	oc.channel <- val
}
