package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

type Request struct {
	File []struct {
		Name     string
		FilePath string
	}
}

func main() {
	var r Request
	r.File = append(r.File, struct {
		Name     string
		FilePath string
	}{Name: "test.apk", FilePath: "media/ceping/98b9cb54-d450-4d66-af39-86bced8fb01d/计算器.apk"})

	out, err := os.Create("test.zip")
	if err != nil {
		fmt.Println("Create error: ", err)
		return
	}

	defer out.Close()

	writer := zip.NewWriter(out)

	for _, file := range r.File {
		fileWriter, err := writer.Create(file.Name)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("权限不足: ", err)
				return
			}
			fmt.Printf("Create file %s error: %s\n", file.Name, err.Error())
			return
		}
		fileInfo, err := os.Open(file.FilePath)
		if err != nil {
			fmt.Printf("Open file %s error: %s\n", file.FilePath, err.Error())
			return
		}
		fileBody, err := ioutil.ReadAll(fileInfo)
		if err != nil {
			fmt.Printf("Read file %s error: %s\n", file.FilePath, err.Error())
			return
		}

		_, err = fileWriter.Write(fileBody)
		if err != nil {
			fmt.Println("Write file error: ", err)
			return
		}
	}

	if err := writer.Close(); err != nil {
		fmt.Println("Close error: ", err)
	}

}
