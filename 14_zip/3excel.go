package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	//f := excelize.NewFile()
	//// Create a new sheet.
	//
	//// Set value of a cell.
	//f.SetCellValue("Sheet1", "A1", 100)
	//f.SetCellValue("Sheet1", "", 100)
	//f.SetCellValue("Sheet1", "A1", 100)
	//f.SetCellValue("Sheet1", "A1", 100)
	//f.SetCellValue("Sheet1", "A1", 100)
	//f.SetCellValue("Sheet1", "A1", 100)
	//f.SetCellValue("Sheet1", "A1", 100)
	//// Set active sheet of the workbook.
	//
	//// Save spreadsheet by the given path.
	//if err := f.SaveAs("Book1.xlsx"); err != nil {
	//	fmt.Println(err)
	//}

	//ReadExcel()
	sss()
}

func ReadExcel() {
	f, err := excelize.OpenFile("/Users/blj/Downloads/skybai/learn/14_zip/7600H1硕软卡.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 读取第一列的数据
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	var countMap = make(map[string]struct{})
	for _, row := range rows {
		for _, colCell := range row {
			i++
			countMap[colCell] = struct{}{}
			fmt.Printf("%s\n", colCell)
		}
	}
	fmt.Println(len(countMap), i)

}

func sss() {
	for i := 0; i < 100; i++ {
		if i < 10 {
			fmt.Println("111", i)
		} else if i < 50 {
			fmt.Println("222", i)

		}
	}
}
