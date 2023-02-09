package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	err := DownloadFile("https://car3.autoimg.cn/cardfs/series/g27/M05/AB/2E/autohomecar__wKgHHls8hiKADrqGAABK67H4HUI503.png", "/Users/blj/Desktop/te")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DownloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("DownloadFile: http.Get err.\n%w", err)
	}
	defer resp.Body.Close()

	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("DownloadFile: os.Create err.\n%w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("DownloadFile: io.Copy err.\n%w", err)
	}
	return nil
}
