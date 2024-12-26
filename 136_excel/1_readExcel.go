package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	excelToMap()
	//excelToMap2(`{"imeis":"1234567890"}`)
}

// excelToMap 查看采集订阅是否重复
func excelToMap() {
	f, err := excelize.OpenFile("/Users/blj/Desktop/采集订阅imei/09042.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0
	arr := []string{}
	for _, row := range rows {
		// 获取第一列的数据
		i++
		if len(row) == 0 {
			continue
		}
		arr = append(arr, row[0])
		//fmt.Printf(row[0])
	}
	fmt.Println("total:", i)
	arr, dupArr, total, duplicates := removeDuplicates(arr)
	fmt.Println("Array after removing duplicates:", len(arr))
	fmt.Printf("Total elements: %d\n", total)
	fmt.Printf("Duplicate elements: %d\n", duplicates)
	fmt.Printf("Duplicate elements: %v\n", dupArr)
	//fmt.Println("sss", arr)

	//arrStr := strings.Join(arr, ",")
	//fmt.Println(arrStr)
	//jarr, _ := json.Marshal(arr)
	//fmt.Println("-------------", string(jarr))
}

func removeDuplicates(elements []string) ([]string, []string, int, int) {
	encountered := map[string]bool{}
	result := []string{}
	duplicateData := []string{}
	total := len(elements)
	duplicates := 0

	for _, element := range elements {
		if !encountered[element] {
			encountered[element] = true
			result = append(result, element)
		} else {
			duplicates++
			duplicateData = append(duplicateData, element)
		}
	}

	return result, duplicateData, total, duplicates
}

// 数据库存data 没有 时分秒 go 用 string去保存 rpc 也用 string
//  startTime 的格式已经是 "2006-01-02 15:04:05"（即标准的日期时间格式）
