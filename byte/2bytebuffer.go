package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte("Learning"))
	buf.WriteString(" Go")
	fmt.Println(buf.String())

}

// 也就是说要有一块地[]byte 去保存我们需要存的东西 还要有从这块地里面读的方法 和 写的方法
