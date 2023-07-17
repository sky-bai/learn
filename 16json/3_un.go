package main

import (
	"encoding/json"
	"fmt"
)

type ConfigBuBiao struct {
	HostsConfig []Hosts `json:"hosts"`
	Signature   string  `json:"signature"`
}

// Hosts 设备对应的部标平台设置
type Hosts struct {
	BkServerPort int    `json:"bk_serverport"`
	BkServerDns  string `json:"bk_serverdns"`
	Jtt808Ver    int    `json:"jtt808ver"`
	ServerPort   int    `json:"serverport"`
	ServerDns    string `json:"serverdns"`
}

func main() {
	configStr := "{\"ImeiStr\":\"803278214323710\",\"ConfigStr\":\"{\\\"hosts\\\":[{\\\"serverdns\\\":\\\"129.204.23.58\\\",\\\"serverport\\\":\\\"7611\\\",\\\"bk_serverdns\\\":\\\"129.204.23.58\\\",\\\"bk_serverport\\\":\\\"7611\\\"}],\\\"signature\\\":\\\"v1.1\\\"}\",\"Scheme\":1,\"Platform\":\"tongtianxin\",\"IsOpen\":false,\"token\":\"GmVf4vH6xMu8NNHeW7VQ\",\"appPlatform\":\"work-space\",\"appVersion\":\"0.1.1\"}"
	var configBuBiao ConfigBuBiao
	err := json.Unmarshal([]byte(configStr), &configBuBiao)
	if err != nil {
		fmt.Println("error:", err)
	}
}
