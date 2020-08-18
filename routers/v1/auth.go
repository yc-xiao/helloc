package v1

import (
	"Helloc/models"
	"Helloc/utils"
	"github.com/gin-gonic/gin"
)
// @Summary 登录认证
// @Description 通过JWT验证账号密码，获取token
// @Tags 认证
// @Accept json
// @Param json body string true "用户id"
// @Success 200 {string} json "{"msg": "token生成成功!", "results": "tokenString"}"
// @Failure 400 {string} json "{"msg": "参数错误/token生成失败!/账号不存在或密码错误", "results": null}"
// @Router /auth/ [post]
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