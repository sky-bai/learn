package main

import qrcode "github.com/skip2/go-qrcode"
import "fmt"

func main() {
	err := qrcode.WriteFile("imei_xxx@sn_xxx", qrcode.Medium, 256, "qr.jpg")
	if err != nil {
		fmt.Println("write error")
	}
}
