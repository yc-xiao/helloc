package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 中间件　-> 返回一个 func(ctx *gin.Context) -> 存到一个array -> next调用下个函数(类似递归)

func TestMiddleWare1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("开始中间件1")
		ctx.Next()
		fmt.Println("结束中间件1")
	}
}

func TestMiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("开始中间件2")
		ctx.Next()
		fmt.Println("结束中间件2")
	}
}

func TestMiddleWare3(ctx *gin.Context) {
	fmt.Println("开始中间件3")
	ctx.Next()
	fmt.Println("结束中间件3")
}