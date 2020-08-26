package middlewares

import (
	"Helloc/models"
	db "Helloc/models/utils"
	"Helloc/utils"
	"fmt"
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
		userId := strconv.Itoa(user.Id)
		if ok && (user.IsAdmin || userId == ctx.Param("id")){

		}else{
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
		}
		log.Println("out OwnerOrAdmin")
	}
}

func OwnerOrAdminToVideo() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		log.Println("in OwnerOrAdminToVideo")
		u, _ := ctx.Get("user")
		user, ok := u.(*utils.UserInfo)
		getSql := fmt.Sprintf("select * from video where id=%s and userId=%d", ctx.Param("id"), user.Id)
		if ok && (user.IsAdmin || db.Get(new(models.Video), getSql)){

		}else{
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
		}
		log.Println("out OwnerOrAdminToVideo")
	}
}

func OwnerOrAdminToComment() (handlerFunc gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		log.Println("in OwnerOrAdminToComment")
		u, _ := ctx.Get("user")
		user, ok := u.(*utils.UserInfo)
		if !ok {

		}
		id := ctx.Param("id")

		fmt.Println(user, ok, id)
		getSql := fmt.Sprintf("select * from comment where id=%s and userId=%d", ctx.Param("id"), user.Id)
		fmt.Println(getSql)
		c := new(models.Comment)
		if ok && (user.IsAdmin || db.Get(c, getSql)){

		}else{
			utils.HttpBadRequest(ctx, "没有权限访问!", nil)
		}
		log.Println("out OwnerOrAdminToComment")
	}
}