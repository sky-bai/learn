package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")

	ctx, _ = context.WithTimeout(context.Background(), 1*time.Second)
	Video2 = fake2Search(ctx, "video2")
)

type Result string // 这是抽象出来的结果

// Search 抽离出需要并发执行函数的抽象层  确定入参 和 结果           根据什么又要获得什么  通用的入参
type Search func(ctx context.Context, query string) (Result, error) // 然后可以抽象出一组任务数组

// 可以定义不同的高阶函数去做不同类型的任务

func fakeSearch(kind string) Search { // 高阶函数写低阶函数具体的业务逻辑,然后返回低阶函数
	return func(_ context.Context, query string) (Result, error) {
		time.Sleep(time.Duration(3) * time.Second)
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

// 高阶函数返回的是姐弟函数的任务抽象， 下面确定一组任务时，就可以传入具体的高阶函数列表， 然后就可以执行这一组任务了,
// 任务数组是低阶函数的抽象，但是传入的是真正执行操作的高阶函数，这样就可以把任务的执行和任务的定义分开了。
// 后面我再看见函数式编程的时候，就可以理解这是一组操作的抽象。

// 抽象出来的结果和低阶函数相联系
// 高阶函数和低阶函数相联系

// 不同高阶函数写不同的业务逻辑

func fake2Search(ctx context.Context, kind string) Search {
	return func(ctx context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s 2222 for %q", kind, query)), nil
	}
}

func main() {
	Google := func(ctx context.Context, query string) ([]Result, error) { // 因为都用到了errgroup,肯定都是一组任务
		g, ctx := errgroup.WithContext(ctx)
		searches := []Search{Web, Image, Video, Video2} // 1.传入具体的高阶函数列表
		results := make([]Result, len(searches))        // 2.确定结果列表
		for i, search := range searches {
			i, search := i, search // https://golang.org/doc/faq#closures_and_goroutines  // 为了解决这个问题, 我们在这里使用了一个技巧, 就是在闭包中使用外部的变量, 但是在闭包中使用的是外部变量的副本, 这样就不会有问题了
			g.Go(func() error {    // 传入的函数是一个闭包, 闭包中的i,search是外部的变量, 会被复制一份, 但是是值传递, 所以会有问题
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	results, err := Google(ctx, "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

// 场景是并发执行一组任务，所以先确定一组任务的抽象层，然后确定入参和结果，最后确定并发执行的函数 万一是不同类型的任务结果该怎么办

// 可是一组任务如何去控制超时呢？
