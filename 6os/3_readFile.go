package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	s, err := ReadFile3("/Users/blj/Downloads/skybai/learn/6os/test14.log")
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return
	}
	// 创建句柄
	//fileHanle, err := os.OpenFile("/Users/blj/Downloads/skybai/learn/6os/test14.log", os.O_RDONLY, 0666)
	//if err != nil {
	//	fmt.Printf("ss", err)
	//}
	//
	//defer fileHanle.Close()
	//line, err := io.ReadAll(fileHanle)
	//if err != nil {
	//	fmt.Printf("ss", err)
	//}
	//// 创建 Reader
	////var maxKey string
	////var maxVal int
	////m := make(map[string]int)
	////mx := sync.Mutex{}
	//i := 0
	//arr := strings.Split(string(line), "\n")
	//for _, v := range arr {
	//
	//	// 只统计 有 STT 的行
	//	if strings.Contains(v, "[STT") {
	//		i++
	//		//fmt.Println(v)
	//		//if strings.Contains(line, "?") {
	//		//	start := strings.Index(line, "/")
	//		//	end := strings.Index(line, "?")
	//		//	api := line[start:end]
	//		//	mx.Lock()
	//		//	m[api]++
	//		//	mx.Unlock()
	//		//	continue
	//		//}
	//		//start := strings.Index(line, "/")
	//		//end := strings.Index(line, "|")
	//		//api := line[start:end]
	//		//mx.Lock()
	//		//m[api]++
	//		//mx.Unlock()
	//
	//	}
	//
	//}
	//fmt.Println("i", i)
	////fmt.Println(m)
	//maxVal = 0
	//for k, v := range m {
	//	if v > maxVal {
	//		maxKey = k
	//		maxVal = v
	//	}
	//}
	//fmt.Println(maxKey, maxVal)
	s1, err := ReadFile3("/Users/blj/Downloads/skybai/learn/6os/test50.log")
	if err != nil {
		fmt.Printf("read file error:%v\n", err)
		return
	}
	// 统计每个元素出现的次数
	counts := make(map[string]int)
	for _, elem := range s {
		counts[elem]++
	}
	for _, elem := range s1 {
		counts[elem]++
	}

	// 将出现次数排序
	type kv struct {
		Key   string
		Value int
	}
	var sortedCounts []kv
	for k, v := range counts {
		sortedCounts = append(sortedCounts, kv{k, v})
	}
	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].Value > sortedCounts[j].Value
	})
	fmt.Printf("统计出一共有%d个接口\n", len(sortedCounts))
	// 打印出现次数前10的元素和次数
	fmt.Println("所有元素和次数为：")
	for i := 0; i < 260 && i < len(sortedCounts); i++ {
		fmt.Printf("%v 出现了 %v 次\n", sortedCounts[i].Key, sortedCounts[i].Value)
	}

	var deviceCounts []kv
	// find设备相关的前20接口
	for i := 0; i < 20 && i < len(sortedCounts); i++ {
		if strings.Contains(sortedCounts[i].Key, "device") || strings.Contains(sortedCounts[i].Key, "vehicle") {
			deviceKv := kv{sortedCounts[i].Key, sortedCounts[i].Value}
			deviceCounts = append(deviceCounts, deviceKv)
		}
	}

	sort.Slice(deviceCounts, func(i, j int) bool {
		return deviceCounts[i].Value > deviceCounts[j].Value
	})
	fmt.Println("下面是与设备相关的接口")
	fmt.Printf("统计出一共有%d个接口\n", len(deviceCounts))
	// 打印出现次数前10的元素和次数
	fmt.Println("所有元素和次数为：")
	for i := 0; i < 260 && i < len(deviceCounts); i++ {
		fmt.Printf("%v 出现了 %v 次\n", deviceCounts[i].Key, deviceCounts[i].Value)
	}
}

func ReadFile3(path string) (y []string, err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer fileHanle.Close()

	scanner := bufio.NewScanner(fileHanle)
	s := make([]string, 0)

	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // sed awk
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "[STT") {
			m := make(map[string]int)

			mx := sync.Mutex{}
			if strings.Contains(line, "/") && strings.Contains(line, "|") {
				if strings.Contains(line, "?") {
					start := strings.Index(line, "/")
					end := strings.Index(line, "?")
					api := line[start:end]
					mx.Lock()
					m[api]++
					mx.Unlock()
					if api == "/" {
						continue
					}
					s = append(s, api)
					continue
				}

				start := strings.Index(line, "/")
				end := strings.Index(line, "|")
				api := line[start:end]
				mx.Lock()
				m[api]++
				mx.Unlock()
				s = append(s, api)
			}
		}
	}

	return s, nil
}
