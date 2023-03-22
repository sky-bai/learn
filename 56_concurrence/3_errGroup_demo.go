package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type ResultNew string
type SearchNew func(ctx context.Context, query string) (ResultNew, error)

func myFunc(kind, query string) (ResultNew, error) {
	// 模拟函数调用，web耗时2秒成功，image耗时4秒失败，video耗时6秒成功
	if kind == "web" {
		fmt.Println("web start")
		time.Sleep(2 * time.Second)
		fmt.Println("web end")
		return ResultNew(fmt.Sprintf("%s result for %q", kind, query)), nil
	} else if kind == "image" {
		fmt.Println("image start")
		time.Sleep(4 * time.Second)
		fmt.Println("image end")
		return "", errors.New("image failed")
	} else {
		fmt.Println("video start")
		time.Sleep(6 * time.Second)
		fmt.Println("video end")
		return ResultNew(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}
func fakeSearchNew(kind string) SearchNew {
	return func(ctx context.Context, query string) (ResultNew, error) {
		done := make(chan ResultNew, 1) // buf chan, 防止ctx直接返回时goroutine阻塞， 用来传递返回的Result
		errch := make(chan error, 1)    // buf chan, 防止ctx直接返回时goroutine阻塞， 用来传递返回的err
		go func() {
			// 添加一个验证是否goroutine泄露的打印，如果有，说明每个goroutine都执行完了，不存在泄露
			defer func() {
				fmt.Printf("%s back goroutine end\n", kind)
			}()
			resp, err := myFunc(kind, query)
			if err != nil {
				errch <- err
				close(done)
			} else {
				done <- resp
			}
		}()
		select {
		case <-ctx.Done(): // 如果这里返回，会造成goroutine泄露，该怎么办？chan buf设置为1
			fmt.Printf("%s ctx.Done\n", kind) // 由于image返回报错，导致kind为video时不等后台myFunc执行完就直接走到这里，
			return "", ctx.Err()
		case rsp, ok := <-done: // 使用done是否关闭来区分返回成功和失败
			if !ok {
				err := <-errch
				return "", err
			} else {
				return rsp, nil
			}
		}
	}
}

var (
	WebNew   = fakeSearchNew("web")
	ImageNew = fakeSearchNew("image")
	VideoNew = fakeSearchNew("video")
)

func main() {
	Google := func(ctx context.Context, query string) ([]ResultNew, error) {
		g, ctx := errgroup.WithContext(ctx)

		searches := []SearchNew{WebNew, ImageNew, VideoNew}
		results := make([]ResultNew, len(searches))

		for i, search := range searches {
			i, search := i, search
			// g.Go只接收无参函数，此无参函数内部，大部分变量都是函数外的变量，因此形成闭包
			g.Go(func() error {
				// search:     是外部变量，且由于延迟绑定，前一步必须重新赋值给search；
				// ctx, query: 是外部变量，但是不会变化，所以使用安全？
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result // 给外面的results赋值？这里不会造成data race吗(存在data race)？
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Printf("Google err: %v\n", err)

		time.Sleep(5 * time.Second)
		return
	}

	time.Sleep(5 * time.Second)
	// 都成功，打印所有结果
	for _, result := range results {
		fmt.Println(result)
	}

}

// errGroup的使用场景
// 就达到了一个任务失败，全部任务级联取消的效果
// 要求数据的完整性，比如一个任务失败，其他任务也级连取消。
