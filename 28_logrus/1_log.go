package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	method2()
	method2()
}

func init() {
	// 日志作为JSON而不是默认的ASCII格式器. 自定义格式信息
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})
	// 输出到标准输出,可以是任何io.Writer
	log.SetOutput(os.Stdout)

	// 只记录xx级别或以上的日志
	log.SetLevel(log.TraceLevel)

	log.SetReportCaller(true)
}

func method2() {
	log.WithFields(log.Fields{
		//"animal": "walrus",
		"size": 10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
