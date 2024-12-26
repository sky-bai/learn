package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var e1 = io.EOF
	var e2 = io.ErrClosedPipe
	var e3 = io.ErrNoProgress
	var e4 = io.ErrShortBuffer
	_, e5 := net.Dial("tcp", "invalid.address:80")
	e6 := os.Remove("/path/to/nonexistent/file")
	var e = errors.Join(e1, e2)
	e = errors.Join(e, e3)
	e = errors.Join(e, e4)
	e = errors.Join(e, e5)
	e = errors.Join(e, e6)
	fmt.Println(e.Error())
	// 输出如下，每一个err一行
	//
	// EOF
	// io: read/write on closed pipe
	// multiple Read calls return no data or error
	// short buffer
	// dial tcp: lookup invalid.address: no such host
	// remove /path/to/nonexistent/file: no such file or directory
	fmt.Println(errors.Unwrap(e)) // nil
	fmt.Println(errors.Is(e, e6)) //true
	fmt.Println(errors.Is(e, e3)) // true
	fmt.Println(errors.Is(e, e1)) // true

}
