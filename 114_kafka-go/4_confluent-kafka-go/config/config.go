package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./connect/config.yaml") // 注意:如果使用相对路径，则是以main.go为当前位置与配置文件之间的路径
	err := viper.ReadInConfig()                  // 查找并读取配置文件
	if err != nil {                              // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	mysql := viper.Get("mysql")
	fmt.Println(mysql)
}
