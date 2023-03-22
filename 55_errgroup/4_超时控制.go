package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	fmt.Println("time", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	g, ctx := errgroup.WithContext(ctx)
	// Use the context with timeout

	defer cancel() // Make sure to cancel the context to release resources
	// Launch a goroutine to do work
	g.Go(func() error {
		// Do some work
		time.Sleep(5 * time.Second)
		return nil
	})
	time.AfterFunc(time.Second, func() {
		cancel()
		ctx.Done()

	})
	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Println("Error time", time.Now())
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("end time", time.Now())
}
