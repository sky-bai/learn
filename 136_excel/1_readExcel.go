package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	//excelToMap()
	excelToMap2(`{"imeis":"1234567890"}`)
}

func excelToMap() {
	f, err := excelize.OpenFile("/Users/blj/Downloads/skybai/learn/136_excel/1k.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("1k")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(rows)
	var imeiMap = make(map[string]struct{})
	i := 0
	for _, row := range rows {
		i++
		dataStr := "\"" + row[0] + "\"" + ":{" + "},"
		fmt.Println(dataStr)
		imeiMap[row[0]] = struct{}{}
	}
	fmt.Println(i)
	//fmt.Println(imeiMap)
	//fmt.Println(len(imeiMap))
}
