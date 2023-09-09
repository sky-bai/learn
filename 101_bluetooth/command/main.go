package main

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func main() {

}

// runBash 执行bash命令并监听错误返回
func runBash(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	//创建获取命令输出管道
	stdout, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return err
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		log.Println("Error:The command is err,", err)
		return err
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				log.Printf("Error :%s\n", err)
				return err
			}
			return nil
		}
		if string(output) != "" {
			return errors.New(string(output))
		}
	}
}
