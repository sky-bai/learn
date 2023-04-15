package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	// https://didicb.glsx.com.cn/rout4-06 14ing&imei=832731909004982&sn=380005689094982&wxId=gh_131d0e3ed6ee
	apiUrl := "https://didicb.glsx.com.cn/router"
	data := url.Values{}
	data.Set("method", "glsx.ddh.access.thirdparty.qrcode")
	data.Set("format", "json")
	data.Set("app_key", "48e5e13229b82c1b4e6e8c96151f0637")
	data.Set("timestamp", "2023-04-06 14:31:29")
	data.Set("deviceType", "6")
	data.Set("sign", "f49f3901c6e9310593536af78a85e634")
	data.Set("v", "1.0.0")
	data.Set("source", "xiaojing")
	data.Set("imei", "832731909004982")
	data.Set("sn", "380005689094982")
	data.Set("wxId", "gh_131d0e3ed6ee")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	replyDiDiHu, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body failed,err:%v\n", err)
	}

	var response DiDiHuResponses
	err = json.Unmarshal([]byte(replyDiDiHu), &response)
	if err != nil {
		fmt.Printf("json Unmarshal failed,err:%v\n", err)
	}
	fmt.Println(response.Url)
	fmt.Println(response.RowUrl)

	str := ""
	err = json.Unmarshal([]byte(str), &response)
	if err != nil {
		fmt.Printf("json Unmarshal failed,err:%v\n", err)
	}
	fmt.Println(response)

}

type DiDiHuResponses struct {
	Url      string `json:"url"`
	RowUrl   string `json:"rowUrl"`
	Message  string `json:"message"`
	Code     string `json:"code"`
	Solution string `json:"solution"`
}
