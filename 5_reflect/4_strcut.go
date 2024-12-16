package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type CustomerBagListReq struct {
	BaseRequest
	ICCID string `json:"iccid,omitempty"` // ICCID参数，选填
}

type BaseRequest struct {
	Method    string `json:"method"`    // 接口名称，例如: sohan.trade.create
	Username  string `json:"username"`  // 用户名，例如: 测试账号
	Timestamp string `json:"timestamp"` // 时间戳，例如: 1461379193134
	Sign      string `json:"sign"`      // 签名，例如: 9A0A8659F005D6984697E
}

func StructToJSONTagMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	// Iterate over fields
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Check if the field has a JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			// Handle embedded structs without JSON tags
			if fieldValue.Kind() == reflect.Struct {
				nestedFields := StructToJSONTagMap(fieldValue.Interface())
				for k, v := range nestedFields {
					result[k] = v
				}
			}
			continue
		}

		// Extract the key from the JSON tag, ignoring options like "omitempty"
		jsonKey := strings.Split(jsonTag, ",")[0]
		result[jsonKey] = fieldValue.Interface()
	}

	return result
}

func main() {
	req := CustomerBagListReq{
		BaseRequest: BaseRequest{
			Method:    "sohan.trade.create",
			Username:  "test_user",
			Timestamp: "1461379193134",
			Sign:      "9A0A8659F005D6984697E",
		},
		ICCID: "12345678901234567890",
	}

	result := StructToJSONTagMap(req)
	// Pretty-print result
	jsonResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonResult))
}
