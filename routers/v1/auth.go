package v1

import (
	"Helloc/models"
	"Helloc/utils"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context)  {
	user := struct {
		Id int `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	u, ok := models.UserCache[user.Id]
	if ok && user.Name == u.Name && user.Password == u.Password {
		token, err := utils.GenerateJwt(u)
		if err != nil {
			utils.HttpBadRequest(ctx, "token生成失败！", nil)
		} else {
			utils.HttpOk(ctx, "token生成成功!", token)
		}
	}else {
		utils.HttpBadRequest(ctx, "账号不存在或密码错误!", nil)
	}

}