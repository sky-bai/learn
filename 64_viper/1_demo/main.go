package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
)

func main() {
	initConfig()
	r := gin.Default()

	r.Run(cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port))
}

var cfg Cfg

type Cfg struct {
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
}
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Log struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func initConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}
	log.Printf("%#v", cfg)
}

func initLog() {
	switch strings.ToLower(cfg.Log.Level) {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	switch strings.ToLower(cfg.Log.Format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
