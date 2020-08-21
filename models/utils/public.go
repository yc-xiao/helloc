package utils

// 定义数据库操作接口

type Model interface {
	New()
	Delete()
	Update()
	Get()
	Filter()
	All()
}


func init() {

}
