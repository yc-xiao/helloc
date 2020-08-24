package models

// 评论(id, objId, model, context, createdTime)
type Comment struct {
	Id int `json:"id" db:"id"` // 评论id
	ObjId int `json:"objId" db:"objId"` // 评论关联对象id
	Model string `json:"model" db:"model"` // 评论关联模型名
	Context string `json:"context" db:"context"` // 内容
	CreatedTime string `json:"createdTime" db:"createdTime"` // 创建时间
}