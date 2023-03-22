package main

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
)

func main() {
	test := &Logon.SearchRequest{
		Query:         "",
		PageNumber:    222,
		ResultPerPage: 111,
		Corpus:        logon.SearchRequest_WEB,
	}
	dst, err := proto.Marshal(test)
	if err != nil {
		return
	}
	fmt.Println(dst)
	fmt.Println(hex.EncodeToString(dst))

	// 解析数据 -- 方式一
	test1 := &logon.SearchRequest{}
	proto.Unmarshal(dst, test1)
	fmt.Printf("方式一 ：unmarshaled message: %v\n", test1)

	// 方式二
	name := "logon.SearchRequest"
	pt := proto.MessageType(name)
	a := reflect.New(pt.Elem()).Interface().(proto.Message)
	proto.Unmarshal(dst, a)
	fmt.Printf("方式二 ：unmarshaled message: %v\n", a)

}

type GpsData struct {
	state protoimpl.MessageState

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`   /**保留6位小数的纬度*/
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"` /**保留6位小数的经度*/
	/**格林威治时间戳*/
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	/**保留1位小数的km/h速度*/
	Speed float32 `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	/**保留1位小数的0-360的方向，0为正北*/
	Bearing float32 `protobuf:"fixed32,5,opt,name=bearing,proto3" json:"direction,omitempty"`
	/**水平gps精度*/
	HAccuracy int32 `protobuf:"fixed32,6,opt,name=haccuracy,proto3" json:"haccuracy,omitempty"`
	/**海拔高度，保留1位小数，单位m*/
	Altitude float32 `protobuf:"fixed32,7,opt,name=altitude,proto3" json:"altitude,omitempty"`
	/**垂直海拔精度*/
	VAccuracy int32 `protobuf:"fixed32,8,opt,name=vaccuracy,proto3" json:"vaccuracy,omitempty"`
	/**卫星数*/
	Satellites uint32 `protobuf:"fixed32,9,opt,name=satellite,proto3" json:"satellites,omitempty"`
}

// 二进制数据
