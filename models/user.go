package models

// (id, nickname, account, password, email, phone, isAdmin, photoFile, createdTime)(微信)
type User struct {
	Id int `json:"id" db:"id" desc:"用户id"`
	NickName string `json:"nickname" db:"nickname" desc:"名称"`
	Account string `json:"account" db:"account" desc:"账号"`
	Password string `json:"password" db:"password" desc:"密码(明文)"`
	Email string `json:"email" db:"email" desc:"邮箱"`
	Phone string `json:"phone" db:"phone" desc:"手机"`
	IsAdmin bool `json:"isAdmin" db:"isAdmin" desc:"管理员"`
	PhotoFile string `json:"photoFile" db:"photoFile" desc:"用户头像(路径)"`
	CreatedTime string `json:"createdTime" db:"createdTime" desc:"创建时间"`
}