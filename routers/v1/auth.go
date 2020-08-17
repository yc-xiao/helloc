package v1

import "github.com/gin-gonic/gin"

func Auth(ctx *gin.Context)  {
	account := struct {
		Account string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	ctx.ShouldBindJSON(&account)

}