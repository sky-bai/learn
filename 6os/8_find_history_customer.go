package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// 下面是应用log，现在需要使用linux命令找到每行有STT的log，并且把appID打印出来输出到一个新的文件中
//2023-07-16 23:35:36.828: [STT - 9AlyKFvpUc][POST] : (193.112.247.55) /apiV2/travel/history | {"sign":"A8D782ABC7604597DFD77AEABF4D03C4","appID":"siweiguangqibentian","imei":"869497051503405","page":"2","pageSize":"10","mobile":"18408237741"}

func main() {
	filePath := flag.String("f", "foo", "filePath")
	flag.Parse()
	fmt.Println(*filePath)
	if *filePath == "foo" {
		fmt.Println("please input file path")
		return
	}
	ReadHistoryCustomer(*filePath)
}

func ReadHistoryCustomer(path string) (err error) {
	// 1.读文件
	fileHandle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	customer := make(map[string]int)

	// 按行处理txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "[STT") && strings.Contains(line, "/apiV2/travel/history") { // 截取每一行第三个引号到第四个引号之间内容的语句 统计重复的行出现的次数 打印出行内容和次数
			// 获取左括号和右括号的位置
			left := strings.Index(line, "{")
			right := strings.Index(line, "}")
			// 截取左括号和右括号之间的内容
			jsonStr := line[left : right+1]
			fmt.Println("------", line)
			var param Param
			// 解析json
			if err := json.Unmarshal([]byte(jsonStr), &param); err != nil {
				fmt.Println("json.Unmarshal failed, err:", err)
			}
			customer[param.AppID]++
		}
	}
	fmt.Println(customer)
	return nil
}

//  下面是应用log 需要用go语言提取appID的值 /apiV2/travel/history | {"appID":"siweiguangqibentian","imei":"823278220965559","page":"1"}

type Param struct {
	AppID string `json:"appID"`
}
