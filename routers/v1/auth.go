package v1

import (
	"Helloc/confs"
	"Helloc/models"
	db "Helloc/models/utils"
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 登录认证
// @Description 通过JWT验证账号密码，获取token
// @Tags 认证
// @Accept json
// @Param body body string true "name(*用户名) password(*密码)" default({"name": "xiaoming", "password": "123456"})
// @Success 200 {string} json "{"msg": "token生成成功!", "results": "tokenString"}"
// @Failure 400 {string} json "{"msg": "参数错误/token生成失败!/账号不存在或密码错误", "results": null}"
// @Router /auth/ [post]
func Auth(ctx *gin.Context)  {
	user := struct {
		Name string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	u := new(models.User)
	authSql := fmt.Sprintf(`select * from user where account="%s" and password="%s";`, user.Name, user.Password)
	if err := db.DB.Get(u, authSql); err != nil {
		log.Println(err)
		utils.HttpBadRequest(ctx, "账号不存在或密码错误!", nil)
	}else{
		token, err := utils.DefaultGenerateJwt(u)
		if err != nil {
			utils.HttpBadRequest(ctx, "token生成失败！", nil)
		} else {
			utils.HttpOk(ctx, "token生成成功!", token)
		}
	}
}

// @Summary 登录认证2，用于Swagger，只做测试
// @Description 用于Swagger，只做测试
// @Tags 认证
// @Accept json
// @Param body body string true "name(*用户名) password(*密码)" default({"name": "xiaoming", "password": "123456"})
// @Success 200 {string} json "{"msg": "token生成成功!", "results": "tokenString"}"
// @Failure 400 {string} json "{"msg": "参数错误/token生成失败!/账号不存在或密码错误", "results": null}"
// @Router /auth2/ [post]
func Auth2(ctx *gin.Context)  {
	user := struct {
		Name string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	u := new(models.User)
	authSql := fmt.Sprintf(`select * from user where account="%s" and password="%s";`, user.Name, user.Password)
	if err := db.DB.Get(u, authSql); err != nil {
		log.Println(err)
		utils.HttpBadRequest(ctx, "账号不存在或密码错误!", nil)
	}else{
		// 测试接口小时后过期
		token, err := utils.GenerateJwt(u, 1)
		if err != nil {
			utils.HttpBadRequest(ctx, "token生成失败！", nil)
		} else {
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			// secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			domain := confs.Cfg["HTTP_PORT"]
			println(domain)
			ctx.SetCookie("token", token, 10*60, "/", domain, false, true)
			utils.HttpOk(ctx, "token生成成功!", token)
		}
	}
}

