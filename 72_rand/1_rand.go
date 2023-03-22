package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"time"
)

func main() {

	intStr := RandomIntStr(10)
	fmt.Println(intStr)
}

func RandomSeconds(max int64) (time.Duration, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return time.Duration(nBig.Int64() * int64(time.Second)), nil
}

// cryptoRandSecure Int returns a uniform random value in [0, max). It panics if max <= 0.
func cryptoRandSecure(min, max int64) (int64, error) {
	if max <= 0 {
		max = 10 // 强制为10位
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return 0, err
	}
	return min + nBig.Int64(), nil
}

// RandomIntStr 随机生成由数字组成的string
func RandomIntStr(length int) string {
	bytes := []byte("0123456789")
	newBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := cryptoRandSecure(0, int64(len(bytes)))
		newBytes[i] = bytes[n]
	}

	return string(newBytes)
}

// RandomStr 所有样本中生成string
func RandomStr(length int) string { // 不加ip
	bytes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	newBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := cryptoRandSecure(0, int64(len(bytes)))
		newBytes[i] = bytes[n]
	}

	return string(newBytes)
}

// RandomCode 去掉了不容易区分的字符
func RandomCode(length int) string {
	// 没有去掉了容易混淆的字符oOLl,9gq,Vv,Uu,I1
	bytes := []byte("ABCDEFGHJKLMNPQRTUVWXYZabcdefhjkmnprtuvwxy134678")

	newBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := cryptoRandSecure(0, int64(len(bytes)))
		newBytes[i] = bytes[n]
	}

	return string(newBytes)
}

// RandomStrWith13Ip 所有样本中生成string 带ip
func RandomStrWith13Ip(length int) string { // 不加ip

	// 从环境变量中获取ip地址
	ip := os.Getenv("GO_IP_ADDRESS")
	ip = "127.0.0.1"
	randLen := 13 - len(ip)

	bytes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	newBytes := make([]byte, randLen+length)
	for i := 0; i < randLen+length; i++ {
		n, _ := cryptoRandSecure(0, int64(length+randLen))
		newBytes[i] = bytes[n]
	}

	return string(newBytes) + ip
}

func other() {
	for i := 0; i < 5; i++ {
		w := RandomStr(16)
		fmt.Println(w)
	}

	fmt.Println("------------")

	//rand.Seed(seedNum)
	for i := 0; i < 5; i++ {
		w := RandomStr(16)
		fmt.Println(w)
	}
	m := make(map[string]int)
	mux := sync.Mutex{}
	for i := 0; i < 30000; i++ {

		go func() {
			w := RandomStr(16)
			mux.Lock()
			if _, ok := m[w]; !ok {
				m[w] = 1
			} else {
				log.Fatal("重复了")
			}
			mux.Unlock()
		}()
	}
	time.Sleep(time.Second * 20)
}
