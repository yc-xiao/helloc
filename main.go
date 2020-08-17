package main

import (
	"Helloc/confs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloc",
		})
	})

	if port, ok := confs.Cfg["HTTP_PORT"]; ok {
		r.Run(port)
	}else {
		r.Run(":9995")
	}
}
