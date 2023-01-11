package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func init() {
	initConfig()
	dynamicConfig()
}

func initConfig() {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigName("base")   // 配置文件名称
	GlobalConfig.AddConfigPath("config") // 从当前目录的哪个文件开始查找
	GlobalConfig.SetConfigType("yaml")   // 配置文件的类型
	err := GlobalConfig.ReadInConfig()   // 读取配置文件
	if err != nil {                      // 可以按照这种写法，处理特定的找不到配置文件的情况
		if v, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(v)
		} else {
			panic(fmt.Errorf("read config err:%s\n", err))
		}
	}
}

// viper支持应用程序在运行中实时读取配置文件的能力。确保在调用 WatchConfig()之前添加所有的configPaths。
func dynamicConfig() {
	GlobalConfig.WatchConfig() // 先编译 再运行    nohub 二进制文件
	GlobalConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("发现配置信息发生变化: %s\n", event.String())
	})
}
