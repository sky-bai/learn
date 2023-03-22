package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	httpClient             = http.Client{}
	baseURL                = "https://funkeinteraktiv.b-cdn.net"
	currentDataEndpoint    = "/current.v4.csv"
	historicalDataEndpoint = "/history.light.v4.csv"
	currentDataURL         = fmt.Sprintf("%s%s", baseURL, currentDataEndpoint)
	historicalDataURL      = fmt.Sprintf("%s%s", baseURL, historicalDataEndpoint)
)

func saveData(ctx context.Context, url, file string) error {
	errChan := make(chan error)

	go func() {
		//
		time.Sleep(6 * time.Second)
		fmt.Println("执行操作")
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			if err != nil {
				return fmt.Errorf("saveData: %w", err)
			}
			return nil
		}
	}

	return nil
}

func saveOriginalData() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	fn := func(ctx context.Context, url, file string) func() error {
		return func() error {
			return saveData(ctx, url, file)
		}
	}

	g.Go(fn(ctx, currentDataURL, "./data/current.csv"))
	g.Go(fn(ctx, historicalDataURL, "./data/historical.csv"))

	g.Go(func() error {
		err := saveData(ctx, currentDataURL, "./data/current.csv")
		return err

	})

	return g.Wait()
}

func main() {
	fmt.Println("time", time.Now())
	err := saveOriginalData()
	if err != nil {
		fmt.Println("time", time.Now())
		fmt.Println(err)
	}
	fmt.Println("time", time.Now())
	time.Sleep(10 * time.Second)
}

// 每一个go方法去异步go业务逻辑，然后再select监听
