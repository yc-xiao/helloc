package models

// (id, userId, title, desc, category, label, rm, duration, videoFile, createdTime, pass)
type Video struct {
	Id int `json:"id" db:"id"` // 视频ID
	UserId int `json:"userId" db:"userId"` // 视频上传用户
	Title string `json:"title" db:"title"` // 标题
	Desc string `json:"desc" db:"desc"` // 描述
	Category string `json:"category" db:"category"` // 视频类型
	Label string `json:"label" db:"label"` // 视频标签，多个自定义标签
	Rm string `json:"rm" db:"rm"` // 视频格式(mp4,flv)
	Duration string `json:"duration" db:"duration"` // 视频时长
	VideoFile string `json:"videoFile" db:"videoFile"` // 视频文件路径
	CreatedTime string `json:"createdTime" db:"createdTime"` // 创建时间
	Pass bool `json:"pass" db:"pass"` // 是否审核通过
}
