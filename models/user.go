package models

// (id, nickname, account, password, email, phone, isAdmin, photoFile, createdTime)(微信)
type User struct {
	Id int `json:"id" db:"id"` // 用户id
	NickName string `json:"nickname" db:"nickname"` //　名称
	Account string `json:"account" db:"account"` // 账号
	Password string `json:"password" db:"password"` //　密码
	Email string `json:"email" db:"email"` // 邮箱
	Phone string `json:"phone" db:"phone"` //　手机
	IsAdmin bool `json:"isAdmin" db:"isAdmin"` // 是否是管理员
	PhotoFile string `json:"photoFile" db:"photoFile"` // 用户头像
	CreatedTime string `json:"createdTime" db:"createdTime"` // 创建时间
}