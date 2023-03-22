package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	fmt.Println("time", time.Now())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// do something
		time.Sleep(3 * time.Second)
		return nil
	})

	if err := g.Wait(); err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Timed out")
		} else {
			log.Println(err)
		}
	}
	fmt.Println("time", time.Now())

}
