package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("6os/test.hdb")
	})
	router.Run(":9090")
}
