package models

// 评论(id, objId, model, context, createdTime)
type Comment struct {
	Id int `json:"id" db:"id" desc:"评论id"`
	ObjId int `json:"objId" db:"objId" desc:"评论关联对象id"`
	Model string `json:"model" db:"model" desc:"评论关联模型名"`
	Context string `json:"context" db:"context" desc:"内容"`
	CreatedTime string `json:"createdTime" db:"createdTime" desc:"创建时间"`
}