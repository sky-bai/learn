package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type param struct {
	Name  string
	Bey   string
	aey   string
	value string
	key   string
}

func main() {
	p1 := make(map[string]any)
	p1["key"] = "field1"
	p1["aey"] = "field3"
	p1["Bey"] = "field4"
	p1["vey"] = "field5"
	p1["tim"] = time.Now()

	var dataDeviceVersion []param
	dataDeviceVersion = append(dataDeviceVersion, param{
		Name: "field1",
	})
	dataDeviceVersion = append(dataDeviceVersion, param{
		Name: "aield1",
	})
	//dataDeviceVersion = append(dataDeviceVersion, p1)
	// 按照字段来排序

	sort.Slice(dataDeviceVersion, func(i, j int) bool {
		return strings.ToLower(dataDeviceVersion[i].Name) < strings.ToLower(dataDeviceVersion[j].Name) // 升序
		//return strings.ToLower(dataDeviceVersion[i].Name) > strings.ToLower(dataDeviceVersion[j].Name) // 降序
	})
	fmt.Println("000", dataDeviceVersion)
	// 遍历map 按照key来排序
	keys := make([]string, 0)
	for k, _ := range p1 {
		keys = append(keys, k)
		//fmt.Println(k)
	}

	s := ""
	sort.Strings(keys)
	fmt.Println("------", keys)
	for k, v := range p1 {
		s = s + k + AnyToString(v)
	}
	fmt.Println("------", s)

}

// AnyToString Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func AnyToString(value any) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
