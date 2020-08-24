package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"` // 返回说明
	Results interface{} `json:"results"`
}

func HttpOk(ctx *gin.Context, msg string, results interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"message": msg, "results": results})
}

func HttpBadRequest(ctx *gin.Context, msg string, results interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": msg, "results": results})
	ctx.Abort()
}