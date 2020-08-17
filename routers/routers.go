package routers

import (
	v "Helloc/routers/v1"
	"github.com/gin-gonic/gin"
)

func AddRouters(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		u := v1.Group("users/")
		{
			u.POST("", v.AddUser)
			u.DELETE(":id", v.DeleteUser)
			u.PUT(":id", v.ModifyUser)
			u.GET(":id", v.GetUser)
			u.GET("", v.GetUsers)
		}
		c := v1.Group("comment/")
		{
			c.POST("", v.AddUser)
			c.DELETE(":id", v.DeleteUser)
			c.PUT(":id", v.ModifyUser)
			c.GET(":id", v.GetUser)
			c.GET("", v.GetUsers)
		}
	}
}