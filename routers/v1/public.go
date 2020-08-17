package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 临时替代数据库

var UserCache = map[int]*User{}

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int `json:"age"`
	Id   int `json:"id" binding:"required"`
}

func Test(c *gin.Context)  {
	fmt.Println("开始test")
	fmt.Println("结束test")
}
