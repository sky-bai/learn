package errgroup

import (
	"context"
)

func fakeRunTask(ctx context.Context) error {
	return nil
}

func ExampleGroup_group() {
	g := KratosGroup{}
	g.Go(fakeRunTask)
	g.Go(fakeRunTask)
	if err := g.Wait(); err != nil {
		// handle err
	}
}

func ExampleGroup_ctx() {
	g := WithContext(context.Background())
	g.Go(fakeRunTask)
	g.Go(fakeRunTask)
	if err := g.Wait(); err != nil {
		// handle err
	}
}

func ExampleGroup_cancel() {
	g := WithCancel(context.Background())
	g.Go(fakeRunTask)
	g.Go(fakeRunTask)
	if err := g.Wait(); err != nil {
		// handle err
	}
}

func ExampleGroup_maxproc() {
	g := KratosGroup{}
	// set max concurrency
	g.GOMAXPROCS(2)
	g.Go(fakeRunTask)
	g.Go(fakeRunTask)
	if err := g.Wait(); err != nil {
		// handle err
	}
}
