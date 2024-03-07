package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	NewSign("Q1hUWDAwMDAxMDAwMzAxMQ37E0843EA3DE19CDFECCF43A9257AFCA")

	appKey := "Q1hUWDAwMDAxMDAwMzAxMQ=="
	appSecret := "37E0843EA3DE19CDFECCF43A9257AFCA"

	var yiKaReqList []YiKaCutFrameStatReq
	var yiKaReq YiKaCutFrameStatReq
	yiKaReq.Iccid = "89860110000000000001" // 卡号
	//num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(result)/1024/1024), 64) // 将流量size字段 单位byte转换成M 保留两位小数
	yiKaReq.FlowValue = 5 // 流量值大小，double类型，单位是M，精确到小数点后两位。 回充流量值的范围 0-1024
	yiKaReq.PayOrderNum = "123"
	yiKaReqList = append(yiKaReqList, yiKaReq)

	data, _ := json.Marshal(yiKaReqList)
	fmt.Println("xxx:", string(data))

	payload := make(map[string]any)
	payload["appKey"] = appKey
	payload["nonce"] = "0hf1mmkge39mjwd9"
	payload["timestamp"] = 1480918492195
	payload["bathIccid"] = string(data)

	//payload["sign"] = GetYiKaSign(payload, appKey, appSecret)
	fmt.Printf("4.sign:%s", GetYiKaSign(payload, appKey, appSecret))
}

func NewSign(str string) {
	base64Str := base64.StdEncoding.EncodeToString(stringToUTF8Bytes(str))

	fmt.Println("xxx111", CreateMd5(base64Str))
}

type YiKaCutFrameStatReq struct {
	Iccid       string  `json:"iccid"`
	FlowValue   float64 `json:"flowValue"`
	PayOrderNum string  `json:"payOrderNum"`
}

func GetYiKaSign(params map[string]any, appKey, appSecret string) string {
	// 循环遍历出key
	var keys []string
	for key, _ := range params {
		keys = append(keys, key)
	}

	// 1.key排序
	sort.Strings(keys)

	// 2.加密值拼接
	str := ""
	for _, key := range keys {
		if key == "sign" {
			continue
		}
		switch params[key].(type) {
		case []interface{}:
			continue
		case map[string]any:
			continue
		}
		if len(str) > 0 {
			str += "&"
		}
		str += key + "=" + AnyToString(params[key])
	}

	fmt.Println("1.拼接之后str:", str)

	// 3.得到的字符串进行Base64编码
	base64Str := base64.StdEncoding.EncodeToString(stringToUTF8Bytes(str))

	fmt.Println("2.base64Str:", base64Str)

	// 4.先将appSecret（应用私钥,由翼卡TSP平台分配)直接拼接得到key(该key由appKey拼接appSecret组成)
	key := appKey + appSecret

	fmt.Println("key:", string(stringToUTF8Bytes(key)))

	// 5.然后将C步骤得到的Base64编码后的字符串用key进行HmacSHA1哈希得到字节数组
	//data := NewHmacSHA1(stringToUTF8Bytes(key), base64Str)

	fmt.Println("key utf ", string(stringToUTF8Bytes(key)))
	fmt.Println("key utf ", string(stringToUTF8Bytes(base64Str)))

	data := NewHmacSHA1(stringToUTF8Bytes(key), stringToUTF8Bytes(base64Str))
	//data := NewHmacSHA1(stringToUTF8Bytes(key), []byte(base64Str))

	// data 转 10进制

	fmt.Println("3.NewHmacSHA1:", data)

	// 6.对D步骤得到的字节数组进行MD5运算得到32位字符串，即为sign
	sign := CreateMd5(data)

	return sign

}

// 将字符串转换为UTF-8编码的字节数组
func stringToUTF8Bytes(s string) []byte {
	var buf bytes.Buffer

	for _, runeValue := range s {
		buf.WriteRune(runeValue)
	}

	return buf.Bytes()
}

func NewHmacSHA1(key, data []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	//io.WriteString(mac, string(data))
	fmt.Println("mac:", mac.Sum(nil)) // mac.Sum(nil) 将mac的hash转成16进制

	data2 := make([]byte, 0)

	data1 := mac.Sum(nil)
	for i := 0; i < len(data1); i++ {
		// 每一位减去256
		data1[i] = byte(int(data1[i]) - 256)
		data2 = append(data2, data1[i])
	}

	fmt.Println("mac1111:", data2) // mac.Sum(nil) 将mac的hash转成16进制
	//fmt.Println("xxxx", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	return hex.EncodeToString(mac.Sum(nil))
}

func CreateMd5(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	result := md5Hash.Sum(nil)
	//fmt.Println("md5:", result)

	//var fixedArray [16]byte
	//for i := 0; i < len(result); i++ {
	// fixedArray[i] = byte(i + 1)
	//}
	//slice := fixedArray[:]
	//fmt.Println("11111", base64.StdEncoding.EncodeToString(slice))
	return hex.EncodeToString(result[:])
}

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
