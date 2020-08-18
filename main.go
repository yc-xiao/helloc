package main

import (
	"Helloc/confs"
	_ "Helloc/docs"
	"Helloc/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// @title HelloC API (Swagger Example)
// @version 1.0
// @description 文档描述略

// @host 127.0.0.1:9995
// @BasePath /api/v1
func main() {
	f, _ := os.Create("HelloC.log")
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routers.AddRouters(r)

	if port, ok := confs.Cfg["HTTP_PORT"]; ok {
		r.Run(port)
	}else {
		r.Run(":9995")
	}
}
