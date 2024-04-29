package main

import (
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"log"
)

//func main() {
//	excelToMap2(`{"imeis":"1234567890"}`)
//}

func excelToMap2(str string) {
	var dd ImeiExportXlsx
	err := json.Unmarshal([]byte(str), &dd)
	if err != nil {
		log.Println("json.Unmarshal error: ", err)
		return
	}
	log.Println("dd: ", dd)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			//logger.Logger.Errorf("countWjDevice Close err:%v", err)
			return
		}
	}()

}

type ImeiExportXlsx struct {
	Imei string `json:"imeis"`
}
