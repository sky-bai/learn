package main

import (
	"io"
	"io/ioutil"
	"os"
)

func Upload(u io.Reader) (err error) {
	// 创建一个临时文件目录
	f, err := ioutil.TempFile("", "upload")
	if err != nil {
		return
	}

	// destroy the file once done
	defer func() {
		m := f.Name()
		f.Close()
		os.Remove(m)
	}()

	// transfer the bytes to the file
	_, err = io.Copy(f, u)
	if err != nil {
		return
	}

	f.Seek(0, 0)

	// process the meta data
	// handle f

	// upload file
	//err = uploadFile(f)
	//if err != nil {
	//	return
	//}
	//
	//return nil

	return nil

}

// 创建一个临时文件目录

// in that 表示因为 in which 表示在哪里
// some kind of 表示某种类型的
// think a as b 表示把a当作b you can think it as a post office
// put a in b

// 交换机 下的 队列
// 建立一个连接 可以建立多个 channel 通道

// 最长子序列 二叉树的遍历

// 只能建立一次连接
