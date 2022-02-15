package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.

	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", 100)
	f.SetCellValue("Sheet1", "", 100)
	f.SetCellValue("Sheet1", "A1", 100)
	f.SetCellValue("Sheet1", "A1", 100)
	f.SetCellValue("Sheet1", "A1", 100)
	f.SetCellValue("Sheet1", "A1", 100)
	f.SetCellValue("Sheet1", "A1", 100)
	// Set active sheet of the workbook.

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
