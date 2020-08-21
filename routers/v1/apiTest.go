package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {
	fmt.Println("开始test")
	a := c.Param("a")
	fmt.Println(a)
	fmt.Println("结束test")
}