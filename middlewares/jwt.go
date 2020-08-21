package middlewares

import (
	"Helloc/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func JwtCheckMiddleWare() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		log.Println("in JwtCheck")
		t := ctx.GetHeader("token")
		if t == "" {
			t2, _ := ctx.Cookie("token")
			t = t2
		}
		_, user, err := utils.CheckJwt(t)
		if err != nil {
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
			return
		}
		ctx.Set("user", user)
		ctx.Next()
		log.Println("out JwtCheck")
	}
}

func IsAdmin() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		log.Println("in IsAdmin")
		u, _ := ctx.Get("user")
		user, ok := u.(*utils.UserInfo)
		if !ok || !user.IsAdmin {
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
		}
		log.Println("out IsAdmin")
	}
}

func OwnerOrAdmin() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		log.Println("in OwnerOrAdmin")
		u, _ := ctx.Get("user")
		user, ok := u.(*utils.UserInfo)
		userId := strconv.Itoa(user.UserId)
		if ok && (user.IsAdmin || userId == ctx.Param("id")){

		}else{
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
		}
		log.Println("out OwnerOrAdmin")
	}
}