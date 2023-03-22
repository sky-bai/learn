package kratos

import (
	"context"
	"testing"
)

type ABC struct {
	CBA int
}

func TestNormal(t *testing.T) {
	var (
		abcs = make(map[int]*ABC)
		g    Group
		err  error
	)
	for i := 0; i < 10; i++ {
		abcs[i] = &ABC{CBA: i}
	}
	g.Go(func(context.Context) (err error) {
		abcs[1].CBA++
		return
	})
	g.Go(func(context.Context) (err error) {
		abcs[2].CBA++
		return
	})
	if err = g.Wait(); err != nil {
		t.Log(err)
	}
	t.Log(abcs)
}

// 什么时候回调用go函数
