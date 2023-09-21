package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/64_viper/util"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		fmt.Println("当前host是: ", util.GlobalConfig.GetString("service.mysql.host"))
		fmt.Println("当前port是: ", util.GlobalConfig.GetString("service.mysql.port"))
		context.JSON(
			200, gin.H{
				"host":   util.GlobalConfig.GetString("service.mysql.host"),
				"11port": util.Cfg.Mysql.Port,
				"redis":  util.GlobalConfig.GetString("service.redis.port"),
			})
	})
	port := util.GlobalConfig.GetString("service.mysql.port")
	port = ":" + port
	err := r.Run(port)
	if err != nil {
		fmt.Println("启动失败: ", err)
		return
	}
}
