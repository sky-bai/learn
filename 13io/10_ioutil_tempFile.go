package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("/", "example")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("---",tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

// 1.生成一个临时文件
// 2.写入内容
// 3.关闭文件
// 4.删除文件

// TempFile在目录dir中创建一个新的临时文件，打开该文件进行读写，
// 并返回结果*os.File。文件名的生成方式是采用pattern并在末尾添加一个随机字符串。
// 如果pattern包含一个"*"，则随机字符串将替换最后一个"*"。
// 如果dir是空字符串，则TempFile使用默认目录存放临时文件(参见os.TempDir)。
// 同时调用TempFile的多个程序不会选择相同的文件。调用者可以使用f.name()来查找文件的路径名。当不再需要该文件时，删除该文件是调用者的责任。
