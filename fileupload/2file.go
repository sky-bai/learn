package fileupload

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func post_data(url string, json string, filename string, filepath string) {
	var buff bytes.Buffer
	// 创建一个Writer
	writer := multipart.NewWriter(&buff)

	// 写入一般的表单字段
	writer.WriteField("json", json)

	// 写入文件字段
	if checkFileIsExist(filepath) {
		// CreateFormFile 第一个参数是表单对应的字段名;第二个字段是对应上传文件的名称
		w, err := writer.CreateFormFile("file", filename)
		if err != nil {
			//创建文件失败
			//checkErr(err)
			return
		}
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			//读取文件发生错误
			//checkErr(err)
			return
		}
		// 把文件内容写入cd
		w.Write(data)
	}

	writer.Close()

	// 发送一个POST请求
	req, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		//创建请求失败
		//checkErr(err)
	}

	// 设置你需要的Header（不要想当然的手动设置Content-Type）
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//req.Header.Set("Content-Type", "text/html;charset=utf-8")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36 Edg/85.0.564.44")

	var client http.Client
	// 执行请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	// 读取返回内容
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//读取失败
		//checkErr(err)
		return
	}

	fmt.Println(string(d))

}

func checkFileIsExist(filepath string) bool {
	var exist = true
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
