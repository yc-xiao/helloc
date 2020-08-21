package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpOk(ctx *gin.Context, msg string, results interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"message": msg, "results": results})
}

func HttpBadRequest(ctx *gin.Context, msg string, results interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": msg, "results": results})
	ctx.Abort()
}