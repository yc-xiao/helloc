package models

// 临时替代数据库

var UserCache = map[int]*User{}

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int `json:"age"`
	Id   int `json:"id" binding:"required"`
}


