package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {
	fmt.Println("开始test")
	fmt.Println("结束test")
}