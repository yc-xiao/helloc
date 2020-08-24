package v1

import (
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TParam struct {
	A string `json:"a" binding:"required" example:"a"` // 字符串a
	B int `json:"b" example:"2"` // 数字b
	C bool `json:"c" example:"false"` // 布尔c
}

type TResponse struct {
	Message string `json:"message"` // 成功返回说明
	Results interface{} `json:"results" `
}

type TResponse2 struct {
	Message string `json:"message"` // 错误返回说明
	Results interface{} `json:"results" `
}

// @Summary 接口测试
// @Description 描述
// @Tags 测试
// @Accept json
// @Param body body TParam false "desc"
// @Success 200 {object} TResponse{results=TParam}
// @Failure 400 {object} TResponse2
// @Router /test/ [post]
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