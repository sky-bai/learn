package main

import (
	"bytes"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
)

func main() {

	//cipherText := "sdfsfddsfsfsdfsfds003fdsfsdfsdfs33dfsdfddsfd222sfddsfdsf33d"
	//key := "888edf9d33fca5dae012f5fa0e5e2079"
	//fmt.Println("len key", len([]byte(key)))
	//plainText, err := sm4.Sm4Cbc([]byte(key),[]byte(cipherText),false)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("plainText:", string(plainText))
	//flag := bytes.Equal(src, plainText)

	src := []byte("这是对称加密SM4的CBC模式加解密测试")
	key := []byte("1q2w3e4r5t6y7u8i")
	cipherText, err := sm4.Sm4Cbc(key, src, true)
	if err != nil {
		panic(err)
	}
	plainText, err := sm4.Sm4Cbc(key, cipherText, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("plainText:", string(plainText))
	flag := bytes.Equal(src, plainText)
	fmt.Println("SM4快速实现加解密，数据填充标准为pksc7，是否解密成功：", flag)
}

// 问题1：不能支持多个文件
