package _9_fileupload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Printf("请求参数：%+v", params)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)

	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func main() {
	param := map[string]interface{}{
		"format":       "mp4",
		"source":       2,
		"videoName":    "9bcf173d930675d1c13696b06475d155",
		"httpProtocol": "HTTP",
	}
	paramMap, err := json.Marshal(param)
	if err != nil {
		panic("json解析错误")
	}
	paramJson := string(paramMap)
	extraParams := map[string]string{
		"username":  "原生****01B20KA04001",
		"token":     "091223******1fab14",
		"password":  "A***!",
		"signature": "9bc********b06475d155",
		"params":    paramJson,
	}
	request, err := newfileUploadRequest("http://api.baidu.com/json/sms/service/VideoUploadService/addVideo", extraParams, "file", "C:\\Users\\Admin\\Desktop\\快手视频\\20210114093809264.mp4")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}
}
