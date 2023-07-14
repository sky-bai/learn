package main

import "strings"

const (
	PositionTypeGps      = 1
	PositionTypeNetwork  = 2
	PositionTypeNoneAuth = 4
)

//  // 定位方式
//  positionType: {
//    gps: 1, // gps定位
//    network: 2, // 网络定位
//    none_auth: 4
//  },

func main() {

}

func ParserToGps(tcpMsg string) TcpGpsMsg {
	s := strings.Split(tcpMsg, ",")
	gpsMsg := TcpGpsMsg{}
	gpsMsg.Protocol = s[0]
	gpsMsg.Imei = strings.Split(s[1], ":")[1]
	gpsMsg.Sequence = s[1]
	gpsMsg.Type = s[2]
	gpsMsg.Ts = s[3]
	gpsMsg.Mcs = s[4]
	gpsMsg.Gps = s[5]
	gpsMsg.Cutframe = s[6]
	return gpsMsg

	//da.Data = []byte("yJEqGqkbHZ<<<#1:862582042075907:1:*,1,GPS,1682329329182,wgs84,40047861:116289569:1682329325182:-782:46:2299:9600:1500:27:0|40047858:116289577:1682329326178:-778:46:2338:9600:1500:27:0|40047855:116289580:1682329327178:-778:46:2414:9600:1500:27:0|40047849:116289576:1682329328178:-778:40:2327:10300:1500:27:0|40047844:116289582:1682329329182:-782:40:2328:10300:1500:27:0,1#")
	// const GPS = [
	//  // {name: 'protocol', pos: 0, separator: ':', idx: 0},
	//  { name: 'imei', pos: 0, separator: ':', idx: 1 },
	//  // {name: 'unknown1', pos: 0, separator: ':', idx: 2},
	//  // {name: 'unknown2', pos: 0, separator: ':', idx: 3},
	//  { name: 'sequence', pos: 1 },
	//  { name: 'type', pos: 2 },
	//  { name: 'ts', pos: 3 },
	//  { name: 'mcs', pos: 4 },    // 坐标系
	//  { name: 'gps', pos: 5 }, // pos 表示以逗号分割的位置
	//  { name: 'cutframe', pos: 6 },   // 数据采集能力
	//];
}

type TcpGpsMsg struct {
	Protocol string `json:"protocol"`
	Imei     string `json:"imei"`
	Sequence string `json:"sequence"`
	Type     string `json:"type"`
	Ts       string `json:"ts"`
	Mcs      string `json:"mcs"`
	Gps      string `json:"gps"`
	Cutframe string `json:"cutframe"`
}

func ParserToXt(tcpMsg string) TcpXtMsg {
	s := strings.Split(tcpMsg, ",")
	xtMsg := TcpXtMsg{}
	xtMsg.Protocol = s[0]
	xtMsg.Imei = strings.Split(s[1], ":")[1]
	xtMsg.Sequence = s[1]
	xtMsg.Type = s[2]
	xtMsg.Position = s[3]
	switch s[4] {
	case "A":
		xtMsg.PositionType = PositionTypeGps
	case "V":
		xtMsg.PositionType = PositionTypeNetwork
	case "D":
		xtMsg.PositionType = PositionTypeNoneAuth
	default:
		xtMsg.PositionType = PositionTypeNoneAuth
	}
	xtMsg.GpsDate = s[5]

	return xtMsg
}

// Xt
// "NGLGILxwBb<<<#1:869497050200318:1:*,0000027D,XT,true+++,V,170504,100819,00000000,00000000,0000,0000,000000010000,5A,4,000064,100#"
// const XT = [
//  { name: 'protocol', pos: 0, separator: ':', idx: 0 },
//  { name: 'imei', pos: 0, separator: ':', idx: 1 },
//  { name: 'unknown1', pos: 0, separator: ':', idx: 2 },
//  { name: 'unknown2', pos: 0, separator: ':', idx: 3 },
//  { name: 'sequence', pos: 1 },
//  { name: 'type', pos: 2 },
//  { name: 'position', pos: 3 },
//  {
//    name: 'positionType',
//    pos: 4,
//    parse: (positionType) => {
//      if (positionType === 'A') {
//        return config.enums.positionType.gps
//      } else if (positionType === 'V') {
//        return config.enums.positionType.network
//      } else if (positionType === 'D') {
//        return config.enums.positionType.none_auth
//      }
//
//      return config.enums.positionType.none_auth
//    }
//  },
//  {
//    name: 'gpsDate',
//    pos: 5,
//    parse: (day, second) => {
//      return new Date(parseInt('0x' + day.substr(0, 2)) + 2000,
//        parseInt('0x' + day.substr(2, 2)) - 1,
//        parseInt('0x' + day.substr(4, 2)),
//        parseInt('0x' + second.substr(0, 2)),
//        parseInt('0x' + second.substr(2, 2)),
//        parseInt('0x' + second.substr(4, 2)));
//    }
//  },
//  { name: 'unknown4', pos: 6 }, // 这个也是和上一个共同组成了gpsDate
//  {
//    name: 'lat',
//    pos: 7,
//    parse: (lat) => {
//      if (lat === '') {
//        return undefined
//      }
//
//      return parseInt('0x' + lat) / 600000 || 0;
//    }
//  },
//  {
//    name: 'lon',
//    pos: 8,
//    parse: (lon) => {
//      if (lon === '') {
//        return undefined
//      }
//
//      return parseInt('0x' + lon) / 600000 || 0;
//    }
//  },
//  { name: 'speed', pos: 9, parse: (speed) => { return parseInt('0x' + speed) / 100 || 0; } },
//  { name: 'direct', pos: 10, parse: (direct) => { return parseInt('0x' + direct) / 100 || 0 } },
//  { name: 'status', pos: 11, parse: (status) => { return getStatus(status) } },
//  // { name: 'uniAlerts', pos: 11, parse: (uniAlerts) => { return getAlerts(uniAlerts) } },
//  { name: 'signal', pos: 12 },
//  { name: 'power', pos: 13 },
//  { name: 'mileage', pos: 14 },
//  { name: 'accuracy', pos: 15, parse: (accuracy) => { return accuracy ? accuracy.split('#')[0] : 50 } } // 相当于每个字段的解析规则
//];

// TcpXtMsg xt 协议
type TcpXtMsg struct {
	Protocol     string `json:"protocol"`
	Imei         string `json:"imei"`
	Unknown1     string `json:"unknown1"`
	Unknown2     string `json:"unknown2"`
	Sequence     string `json:"sequence"`
	Type         string `json:"type"`
	Position     string `json:"position"`
	PositionType int    `json:"positionType"`
	GpsDate      string `json:"gpsDate"`
	Unknown4     string `json:"unknown4"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	Speed        string `json:"speed"`
	Direct       string `json:"direct"`
	Status       string `json:"status"`
	Signal       string `json:"signal"`
	Power        string `json:"power"`
	Mileage      string `json:"mileage"`
	Accuracy     string `json:"accuracy"`
}

// 在Go语言中，可以使用`time.Parse`函数来解析一个给定格式的日期字符串，从而创建一个`time.Time`类型的对象，表示该日期。
//
//以下是一个示例代码，该代码定义了一个名为`ParseDate`的函数，该函数接受两个参数`day`和`second`，这些参数表示16进制的日期和时间字符串。该函数使用`time.Parse`函数将这些字符串解析为`time.Time`对象，并返回该对象。
//
//```go
//import "time"
//
//func ParseDate(day string, second string) (time.Time, error) {
//    layout := "2006-01-02 15:04:05" // 定义日期时间格式
//    dayStr := hexToDec(day)         // 将16进制字符串转换为10进制字符串
//    secondStr := hexToDec(second)
//    dateStr := dayStr + " " + secondStr
//    return time.Parse(layout, dateStr) // 解析日期时间字符串并返回time.Time对象
//}
//
//// 将16进制字符串转换为10进制字符串
//func hexToDec(hex string) string {
//    n, _ := strconv.ParseInt(hex, 16, 64)
//    return strconv.FormatInt(n, 10)
//}
//```
//
//在上述代码中，`layout`变量定义了日期时间字符串的格式，具体格式的说明可以参考Go语言的`time`包文档。然后，`hexToDec`函数将16进制的日期和时间字符串转换为10进制字符串，`dayStr`和`secondStr`变量分别表示转换后的日期和时间字符串。最后，这两个字符串被拼接成一个完整的日期时间字符串`dateStr`，并使用`time.Parse`函数解析为`time.Time`类型的对象，该对象作为函数的返回值。
