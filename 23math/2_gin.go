package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/v3/a/action/liveCallBack", getting)
	//监听端口默认为8080
	r.Run(":8000")

}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "hello word")
}
