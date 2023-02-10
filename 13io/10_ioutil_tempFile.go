package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {

	GetLastE()

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

	err = OsCreateTemp()
	if err != nil {
		panic(err)
	}

	// 创建临时目录
	OsCreateTempDirectory()
}

// 1.生成一个临时文件 磁盘上的 就不用ram 内存了
// 2.写入内容
// 3.关闭文件
// 4.删除文件

// 针对于大文件上传

// TempFile在目录dir中创建一个新的临时文件，打开该文件进行读写，
// 并返回结果*os.File。文件名的生成方式是采用pattern并在末尾添加一个随机字符串。
// 如果pattern包含一个"*"，则随机字符串将替换最后一个"*"。
// 如果dir是空字符串，则TempFile使用默认目录存放临时文件(参见os.TempDir)。
// 同时调用TempFile的多个程序不会选择相同的文件。调用者可以使用f.name()来查找文件的路径名。当不再需要该文件时，删除该文件是调用者的责任。

// OsCreateTemp 当您编写一个处理大量数据的程序，或者您测试一个创建文件的程序时，您通常需要能够创建临时文件，以允许您在短时间内存储数据而不会使项目的磁盘空间混乱。 比如我抓取别人的数据转存到自己的oss上面
func OsCreateTemp() error {
	// os 创建一个临时文件
	f, err := os.CreateTemp("", "example") // 文件名就是 pattern + random string 临时文件
	if err != nil {
		return err
	}
	defer os.Remove(f.Name()) // clean up
	// 如果不删除 会在磁盘上生成一个临时文件 但是不会在内存中生成一个临时文件
	// 临时文件的内容是在内存中的 但是临时文件是在磁盘上的 这样为什么可以避免内存溢出呢？
	// 内存溢出是因为内存中的数据太多了 但是临时文件是在磁盘上的 所以不会造成内存溢出

	if _, err := f.Write([]byte("content")); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// OsCreateTempDirectory 创建临时目录
func OsCreateTempDirectory() {
	// create a temporary directory
	name, err := os.MkdirTemp("", "dir") // in Go version older than 1.17 you can use ioutil.TempDir
	if err != nil {
		log.Fatal(err)
	}

	// remove the temporary directory at the end of the program
	defer os.RemoveAll(name)

	// print path of the directory
	fmt.Println(name)
}

func GetLastE() {
	st := path.Base("D:/go.txt") // 返回路径的最后一个元素 go.txt
	fmt.Println("---", st)
}

// os.Create  os.CreateTemp 三者的区别
// os.Create 创建一个文件 如果文件存在 则清空文件内容 如果文件不存在 则创建一个新的文件
// os.CreateTemp 创建一个临时文件 临时文件是在磁盘上的 但是临时文件的内容是在内存中的
