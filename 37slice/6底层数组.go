package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	a := []int{3, 2, 3, 4, 5}
	fmt.Println("--------", a[:2])
	fmt.Println("time.now", strconv.Itoa(time.Now().Year())[:1])

	fmt.Printf("%+v\n", a)

	ap(a)
	fmt.Printf("%+v\n", a)

	app(a)
	fmt.Printf("%+v\n", a)

	PrintData()
	data()

}

func ap(a []int) {
	a = append(a, 6)
	fmt.Printf("3333%+v\n", a)

}

func app(a []int) {
	a[0] = 1
	fmt.Printf("1111%+v\n", a)

}

// 不会影响原切片长度
// 但是修改值会影响原切片值

func PrintData() {
	xt := "#1:803278214323710:1:*,0000057C,XT,true+四川省+成都市+武侯区,V,170610,0D230F,011798A3,03B8BA1B,01A9,8B74,000000010000,5A,4,000064,6#"
	status := "#1:803278214323710:1:*,000000EA,STATUS,28,57140,3,1,1,0,0,0,0,0,0,0,0.000000,-95,1,-1,-1,0,258#"
	accOff := "#1:0:1:*,803278214323710,ACCOFFMODEL,alarmaccoff#"
	gps := "#1:803278214323710:1:*,1,GPS,1686894543575,wgs84,30541797:104061056:1686894539757:-469:43:43291:1000:1432:25:0|30541817:104061053:1686894540575:-287:42:43603:1700:1184:25:0|30541842:104061051:1686894541575:-287:41:44085:3000:1035:25:0|30541860:104061050:1686894542575:-287:38:44458:3000:931:25:0|30541874:104061047:1686894543575:-287:38:44672:3000:861:25:0,1#"
	up := "#1:803276194700212:1:*,IZDUKXnqpW,UP,MA%3D%3D#"
	cloudState := "1:0:1:*,a5Ax7X7UBaMZ,800277201100050,CLOUDSTATE,0,1,nothing"

	fmt.Println("len:", len(xt+status+accOff+gps+up+cloudState))
}

// 代码
// 每隔30s，就打印

func data() {
	i := 0
	for {
		i++
		if i%30 == 0 {
			fmt.Println("30s")
		}
		time.Sleep(1 * time.Second)
	}

}
