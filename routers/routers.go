package routers

import (
	"Helloc/middlewares"
	v "Helloc/routers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRouters(r *gin.Engine) {
	r.Static("static/", "./static/")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently,"/swagger/index.html")
	})
	v1 := r.Group("api/v1")
	// 登录认证
	{
		v1.POST("auth/", v.Auth)
		v1.POST("auth2/", v.Auth2)
	}
	// 用户
	{
		u := v1.Group("users/")
		u.GET("", v.GetUsers)
		u.Use(middlewares.JwtCheckMiddleWare())
		{
			u.POST("", middlewares.IsAdmin(), v.AddUser)
			u.DELETE(":id", middlewares.IsAdmin(), v.DeleteUser)
			u.PUT(":id", middlewares.OwnerOrAdmin(), v.ModifyUser)
			u.GET(":id", v.GetUser)
			u.PUT(":id/photo", middlewares.OwnerOrAdmin(), v.UploadPhoto)
		}
	}
	// 视频
	{
		// 1.新增视频信息 2.上传视频　3.获取视频详情 4.下载视频
		video := v1.Group("videos/")
		video.GET(":id", v.GetVideo)
		video.GET("", v.GetVideos)

		video.Use(middlewares.JwtCheckMiddleWare())
		{
			video.POST("", v.AddVideo)
			video.DELETE(":id", middlewares.OwnerOrAdminToVideo(), v.DeleteVideo)
			video.PUT(":id", middlewares.OwnerOrAdminToVideo(), v.ModifyVideo)
			video.POST(":id/upload/", middlewares.OwnerOrAdminToVideo(), v.UploadVideo)
			video.POST(":id/check/", middlewares.IsAdmin(), v.CheckVideo)
		}
	}
	// 评论
	{
		comment := r.Group("comments/")
		comment.GET(":id/")
		comment.POST(":id/", middlewares.JwtCheckMiddleWare())
		comment.DELETE(":id/", middlewares.OwnerOrAdminToComment())
	}

	// 测试
	{
		t := v1.Group("test/")
		t.Use(middlewares.TestMiddleWare1())
		t.Use(middlewares.TestMiddleWare3)
		t.GET("", v.Test, middlewares.TestMiddleWare2())
		t.POST("", v.Test, middlewares.TestMiddleWare2())
		// t.Use(middlewares.TestMiddleWare3) 无法在最后添加
	}
}