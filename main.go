package main

import (
	"Helloc/confs"
	"Helloc/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("helloc.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	routers.AddRouters(r)

	if port, ok := confs.Cfg["HTTP_PORT"]; ok {
		r.Run(port)
	}else {
		r.Run(":9995")
	}
}
