package utils

import "strconv"

// 临时替代数据库
var TempUserCache = map[int]*TempUser{}

type TempUser struct {
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
		TempUserCache[i] = &TempUser{i, strconv.Itoa(i),strconv.Itoa(i), isAdmin }
	}
}