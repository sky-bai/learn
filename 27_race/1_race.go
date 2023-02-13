package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	fmt.Println("a is ", a)

	time.Sleep(2 * time.Second)
}

// go run --race 1_race.go

// --go_out有两层含义，一层是输出的是go语言对应的文件；一层是指定生成的go文件的存放位置。 把文件生成在对应目录下
// --go_opt=paths=source_relative 指定生成的文件的路径是相对于proto文件的路径
// --go_opt表示生成go文件时候的目录选项，如上面写时表示生成的文件与proto在同一目录。

// option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";
// go_packge有两层意思，一层是表明如果要引用这个proto生成的文件的时候import后面的路径；一层是如果不指定--go_opt（默认值），生成的go文件存放的路径。
// go_package的含义是生成的go文件的包名，如果不指定--go_opt（默认值），生成的go文件存放的路径。

// proto 文件中的package package用于防止不同project之间定义了同名message结构的冲突，如果不指定package，那么默认的package就是当前文件所在的目录名。
// 所以package的命名一般是按照目录名来命名的，比如当前文件所在的目录名是tutorial，那么package就是tutorial。 文件目录就是功能
// package 是跟到 proto文件走的
