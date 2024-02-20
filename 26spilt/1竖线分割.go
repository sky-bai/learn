package main

import (
	"fmt"
	"strings"
)

func main() {
	//strllll := "中心领导|开发处室领导|测试处室领导|开发人员|测试人员"
	//st2 := "中心领导|开发处室领导|"
	//sr := strings.Split("中心领导|开发处室领导|测试处室领导|开发人员|测试人员", "|")
	//fmt.Println("sr:", sr)
	//is1 := strings.Contains(strllll, "中心领导")
	//fmt.Println("is1:", is1)
	//
	//if strings.Contains(strllll, "开发处室领导") || strings.Contains(strllll, "测试处室领导") {
	//	fmt.Println("该角色是部门管理员")
	//}
	//if strings.Contains(st2, "开发处室领导") || strings.Contains(st2, "测试处室领导") {
	//	fmt.Println("该角色是部门管理员")
	//}

	// #1:{imei}:1:*,{sequence},GPS,{ts},{gcs},{gps},{cutframe}#
	msg := "#1:803278214323702:1:*,1,GPS,1705462251742,wgs84,30541697:104061040:1705462247688:-688:0:47028:18700:545:9:0:3|30541697:104061040:1705462248687:-687:0:47028:18100:545:9:0:3|30541500:104061033:1705462249687:-687:0:47663:800:546:9:0:3|30541587:104061050:1705462250687:-687:0:48293:900:546:9:0:3|30541587:104061050:1705462251688:-688:0:48293:900:546:9:0:3,3#\n[2024-01-17 11:30:55] [info] [uUQCNAUuAd] < #1:803278214323702:1:*,0000001B,XT,true+四川省+成都市+武侯区,A,180111,0B1E37,011797B7,03B8BAE8,0000,0320,000000010000,5A,4,000064,5#"
	arr := strings.Split(msg, ",")

	fmt.Println("arr:", len(arr))

	if len(arr) > 6 {
		fmt.Println("arr[6]:", arr[6])
	}

	// http://sdktest.e-car.cn:7070/service/pflow/renew/increaseFlow?appKey=testD1ECEAxAQ13shpYR&bathIccid=%5B%7B%22iccid%22%3A%2289860619120050967116%22%2C%22flowValue%22%3A5.03%2C%22payOrderNum%22%3A%221705480212780%22%7D%5D&nonce=1ws7EQFrdC&sign=e5e6c7a62e2ec2a9c8c2af250ddf88dc&timestamp=1705480212780

}

func is1(s string) bool {
	return strings.Contains(s, "1")
}
