package middlewares

import (
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JwtCheckMiddleWare() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		fmt.Println("jwt check")
		t := ctx.GetHeader("token")
		_, user, err := utils.CheckJwt(t)
		if err != nil {
			msg := fmt.Sprintf("token认证失败, 错误原因 %s", err)
			utils.HttpBadRequest(ctx, msg, nil)
			ctx.Abort()
		}
		ctx.Set("user", user)
		ctx.Next()
		fmt.Println("jwt check end")
	}
}

func IsAdmin() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		fmt.Println("admin")
		u, _ := ctx.Get("user")
		user, ok := u.(*utils.UserInfo)
		if !ok || !user.IsAdmin {
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
			ctx.Abort()
		}
		fmt.Println("admin close")
	}
}