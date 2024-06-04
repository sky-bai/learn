package main

import (
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	// 给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

	// 读取PEM文件
	pemData, err := os.ReadFile("/Users/blj/Desktop/temp/apple_g3_root/apple_root.pem")
	if err != nil {
		log.Fatalf("无法读取PEM文件: %v", err)
	}

	// 解析PEM块
	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "CERTIFICATE" {
		log.Fatalf("未找到有效的证书PEM数据")
	}

	// 获取PEM编码的证书内容
	pemEncodedCertificate := pem.EncodeToMemory(block)

	// 打印证书内容
	fmt.Printf("PEM编码的证书内容:\n%s\n", pemEncodedCertificate)

	// 打印文件中剩余的内容（如果有）
	if len(rest) > 0 {
		fmt.Printf("文件中剩余的内容:\n%s\n", rest)
	}
}
