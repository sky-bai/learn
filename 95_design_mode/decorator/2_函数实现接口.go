package main

import "fmt"

type Job interface {
	Run()
}

type FuncJob func()

// 定义函数类型

func (f FuncJob) Run() {
	f()
}

func main() {
	// 定义不同类型的任务函数
	task1 := FuncJob(func() {
		fmt.Println("Running task 1")
	})
	// 创建具体对象

	task2 := FuncJob(func() {
		fmt.Println("Running task 2")
	})

	// 使用任务调度器执行任务
	jobScheduler(task1)
	jobScheduler(task2)
}

func jobScheduler(job Job) {
	// 调度器执行任务
	job.Run()
}

// 一般接口就是作为入参和返回值 作为函数入参的时候，直接调用接口里面的方法
// 作为返回值的时候 那么我就要传入一个实现了这个接口的对象,不管是函数类型还是其他结构体类型。

// 对于函数实现了接口 我们就想想该如何使用

// 使用的时候调用接口的具体方法进行调用

// 当说到函数类型实现接口的灵活性时，意味着我们可以轻松地将不同的函数逻辑作为对象进行传递，并使用它们来满足相同接口的需求。这样做可以实现代码的解耦和扩展，让我们能够更灵活地处理不同的行为和功能。
//
//以下是一个例子，说明了函数类型作为接口方法的灵活性。
//
//假设我们有一个简单的任务调度器，它可以调度不同类型的任务并执行它们。我们定义了一个 `Job` 接口，它具有一个 `Run()` 方法用于执行任务。然后，我们可以使用不同的函数来实现 `Job` 接口的 `Run()` 方法，以执行不同类型的任务。
//
//在这个例子中，我们定义了两个任务函数 `task1` 和 `task2`，它们都是 `FuncJob` 类型，实现了 `Job` 接口的 `Run()` 方法。每个任务函数具有自己的逻辑，例如输出特定的消息。
//
//我们可以将这些任务函数传递给任务调度器 `jobScheduler`，并执行它们。调度器只需要知道接口 `Job` 的定义，而不需要关心具体的任务逻辑。这样，我们可以轻松地添加新的任务函数，并根据需要调度执行它们。
//
//通过使用函数类型实现接口的灵活性，我们可以轻松地定义不同的任务逻辑，并将它们传递给调度器或其他接口需要的地方。这样，我们可以根据需求动态地配置和扩展代码的行为，而不需要修改现有的调度器或接口实现。这种灵活性使得代码更容易维护、扩展和重用。
