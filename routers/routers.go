package routers

import (
	"Helloc/middlewares"
	v "Helloc/routers/v1"
	"github.com/gin-gonic/gin"
)

func AddRouters(r *gin.Engine) {
	r.Static("static/", "./static/")

	v1 := r.Group("api/v1")
	{
		v1.POST("auth/", v.Auth)
	}

	{
		u := v1.Group("users/")
		{
			u.POST("", v.AddUser)
			u.DELETE(":id", v.DeleteUser)
			u.PUT(":id", v.ModifyUser)
			u.GET(":id", v.GetUser)
			u.GET("", v.GetUsers)
		}
	}

	{
		t := v1.Group("test/")
		t.Use(middlewares.TestMiddleWare1())
		t.Use(middlewares.TestMiddleWare3)
		t.GET("", v.Test, middlewares.TestMiddleWare2())
		// t.Use(middlewares.TestMiddleWare3) 无法在最后添加
	}
}