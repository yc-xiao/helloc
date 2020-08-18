package models

import "strconv"

// 临时替代数据库

var UserCache = map[int]*User{}

type User struct {
	Id   int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool `json:"isAdmin"`
}

func init() {
	var isAdmin bool
	for i:=0; i<=120; i++ {
		if i % 3 == 2 {
			isAdmin = true
		}else{
			isAdmin = false
		}
		UserCache[i] = &User{i, strconv.Itoa(i),strconv.Itoa(i), isAdmin }
	}
}

