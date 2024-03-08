package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	//vp.SetConfigName("config")
	//vp.AddConfigPath("configs/")
	//vp.SetConfigType("yaml")

	vp.SetConfigType("yaml")
	vp.SetConfigFile("/Users/blj/Downloads/skybai/learn/114_kafka-go/4_confluent-kafka-go/config/config.yaml") // 注意:如果使用相对路径，则是以main.go为当前位置与配置文件之间的路径
	err := vp.ReadInConfig()                                                                                   // 查找并读取配置文件
	if err != nil {                                                                                            // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	return &Config{vp}, nil
}

func (s *Config) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
