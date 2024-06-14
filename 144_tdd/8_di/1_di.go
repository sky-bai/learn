package __di

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// 只关心上层实现 不关心具体实现
