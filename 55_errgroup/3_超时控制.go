package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

type ResultTimeout string
type SearchTimeout func(ctx context.Context, query string) (ResultTimeout, error)

func fake2SearchTimeout(kind string) SearchTimeout {
	time.Sleep(3 * time.Second)
	return func(_ context.Context, query string) (ResultTimeout, error) {
		return ResultTimeout(fmt.Sprintf("%s 2222 for %q", kind, query)), nil
	}
}

func main() {
	fmt.Println("000000", time.Now())
	Google := func(ctx context.Context, query string) ([]ResultTimeout, error) {
		g, ctx := errgroup.WithContext(ctx)
		searches := []SearchTimeout{fake2SearchTimeout("Web"), fake2SearchTimeout("Image"), fake2SearchTimeout("Video"), fake2SearchTimeout("Video2")}
		results := make([]ResultTimeout, len(searches))
		for i, search := range searches {
			i, search := i, search
			g.Go(func() error {
				select {
				case <-ctx.Done(): // 监听 context.Done() 信号
					return ctx.Err()
				default:
					result, err := search(ctx, query)
					if err == nil {
						results[i] = result
					}
					return err
				}
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 在函数退出时调用 cancel()
	results, err := Google(ctx, "golang")
	if err != nil {
		fmt.Println("000000", time.Now())
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("000000", time.Now())
}
