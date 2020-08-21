package models

// (id, userId, title, desc, category, label, rm, duration, videoFile, createdTime, pass)
type Video struct {
	Id int `json:"id" db:"id" desc:"视频ID"`
	UserId int `json:"userId" db:"userId" desc:"视频被用户上传"`
	Title string `json:"title" db:"title" desc:"标题"`
	Desc string `json:"desc" db:"desc" desc:"描述"`
	Category string `json:"category" db:"category" desc:"视频类型"`
	Label string `json:"label" db:"label" desc:"视频标签，多个自定义标签"`
	Rm string `json:"rm" db:"rm" desc:"视频格式(mp4,flv)"`
	Duration string `json:"duration" db:"duration" desc:"视频时长"`
	VideoFile string `json:"videoFile" db:"videoFile" desc:"视频文件路径"`
	CreatedTime string `json:"createdTime" db:"createdTime" desc:"创建时间"`
	Pass bool `json:"pass" db:"pass" desc:"是否审核通过"`
}
