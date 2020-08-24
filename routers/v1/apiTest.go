package v1

import (
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TParam struct {
	A string `json:"a" binding:"required" example:"a"`
	B int `json:"b" example:"2"`
	C bool `json:"c" example:"false"`
}

func Test(ctx *gin.Context)  {
	fmt.Println("开始test")
	p := new(TParam)
	if err := ctx.ShouldBindJSON(p); err != nil{
		fmt.Println(err)
		utils.HttpBadRequest(ctx, "错误", nil)
	}else{
		utils.HttpOk(ctx, "成功", p)
	}
	fmt.Println("结束test")
}