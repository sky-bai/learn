package kratos

import (
	"context"
	"fmt"
)

func fakeRunTask(ctx context.Context) error {
	return nil
}

func ExampleGroup_group() {
	g := Group{}
	g.Go(fakeRunTask)
	g.Go(fakeRunTask)
	if err := g.Wait(); err != nil {
		// handle err
		fmt.Println("------", err)
	}
}
