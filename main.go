package main

import (
	"Helloc/confs"
	_ "Helloc/docs"
	"Helloc/models/utils"
	"Helloc/routers"
	"Helloc/tasks"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	defer utils.DB.Close()
	defer utils.RedisClose()
	defer tasks.Close()
	go tasks.StartTasks()

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routers.AddRouters(r)

	if port, ok := confs.Cfg["HTTP_PORT"]; ok {
		r.Run(port)
	}else {
		r.Run(":9995")
	}
}
