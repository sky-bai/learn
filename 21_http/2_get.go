package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	localFile := "/Users/blj/Desktop/" + RandomStr(16)
	RandomStr(16)
	err := DownloadFile("https://car3.autoimg.cn/cardfs/series/g27/M05/AB/2E/autohomecar__wKgHHls8hiKADrqGAABK67H4HUI503.png", localFile)
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

func RandomStr(length int) string {
	bytes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	newBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		newBytes[i] = bytes[RandomInt(0, len(bytes))]
	}

	return string(newBytes)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
