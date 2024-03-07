package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, err := ReadFile3("/Users/blj/Downloads/skybai/learn/6_os/countRetry.log")
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return
	}
}

// ReadFileCount 统计retry次数
func ReadFileCount(path string) (y []string, err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)
	i := 0
	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // sed awk
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "[STT") {

			if strings.Contains(line, "/") && strings.Contains(line, "|") {
				if strings.Contains(line, "?") {

					continue
				}

				start := strings.Index(line, "/")
				end := strings.Index(line, "|")

			}
		}
	}

	return s, nil
}
