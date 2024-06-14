package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestLeakWithGoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	NewLeak(make(chan struct{})).leak()
}

// go test -v -run ^TestLeakWithGoleak$ 里面是需要测试的函数名
