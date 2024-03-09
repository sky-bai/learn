package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	//vp.SetConfigName("config")
	//vp.AddConfigPath("configs/")
	//vp.SetConfigType("yaml")

	filepath := "/Users/blj/Downloads/skybai/learn/114_kafka-go/4_confluent-kafka-go/config/config.yaml"

	vp.SetConfigType("yaml")
	vp.SetConfigFile(filepath) // 注意:如果使用相对路径，则是以main.go为当前位置与配置文件之间的路径

	// 获取文件MD5
	confMD5, err := ReadFileMd5(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// 读取配置文件
	err = vp.ReadInConfig() // 查找并读取配置文件
	if err != nil {         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 设置监控文件
	vp.WatchConfig()

	// 设置配置文件修改回调
	vp.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		tconfMD5, err := ReadFileMd5(filepath)
		if err != nil {
			fmt.Println("ReadFileMd5 err:", err)
		}
		// 比对当前MD5与之前是否相同
		if tconfMD5 == confMD5 {
			return
		}
		// 这说明文件发生了改变.
		confMD5 = tconfMD5

		//err = vp.Unmarshal(&conf)
		//if err != nil {
		//	log.Fatalf("vp.Unmarshal err: %v", err)
		//}

		log.Println("Config file changed!")
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

var confMD5 string

func GetMD5(s []byte) string {
	m := md5.New()
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}

func ReadFileMd5(sfile string) (string, error) {
	ssconfig, err := os.ReadFile(sfile)
	if err != nil {
		return "", err
	}
	return GetMD5(ssconfig), nil
}
