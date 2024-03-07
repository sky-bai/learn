package main

import (
	"fmt"
	"time"
)

func main() {
	tsp := time.Now().UnixMilli() / 1000
	LostTime := time.Unix(tsp, 0).Add(7 * 24 * 3600 * time.Second).Format("2006-01-02 15:04:05")
	fmt.Println(LostTime)

	NonceLostTime := time.Unix(tsp, 0).Add(3 * 24 * 3600 * time.Second).Format("2006-01-02 15:04:05")
	fmt.Println("result.NonceLostTime:", NonceLostTime)
}
