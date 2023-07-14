package main

import (
	"context"
	"fmt"
	"time"
)

type MyContext struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

func main() {

	var myCtx MyContext
	myCtx.Ctx, myCtx.Cancel = context.WithCancel(context.Background())
	i := 0

	go func(ctx MyContext) {
		for {
			select {
			case <-ctx.Ctx.Done():
				fmt.Println("子goroutine1 ctx.Done()", time.Now())
				return
			default:

				select {
				case <-time.After(5 * time.Second):
					fmt.Println("i", i)
				}
			}
		}
	}(myCtx)

	go func(ctx MyContext) {
		for {
			select {
			case <-ctx.Ctx.Done():
				fmt.Println("子goroutine2 ctx.Done()", time.Now())
				return
			case <-time.After(20 * time.Second):
				fmt.Println(" default", time.Now())
				return
			}
		}
	}(myCtx)

	go func(ctx MyContext) {
		time.Sleep(1 * time.Second)
		ctx.Cancel()
		i = 1
		fmt.Println("执行 cancel", time.Now())

	}(myCtx)

	select {}
}
