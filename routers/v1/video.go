package v1

import (
	"Helloc/confs"
	"Helloc/models"
	db "Helloc/models/utils"
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"strings"
)

type AddVideoParam struct {
	UserId int `json:"userId" db:"userId" binding:"required"` // 视频上传用户
	Title string `json:"title" db:"title" binding:"required"` // 标题
	Desc string `json:"desc" db:"desc"` // 描述
	Category string `json:"category" db:"category"` // 视频类型
	Label string `json:"label" db:"label"` // 视频标签，多个自定义标签
	Rm string `json:"rm" db:"rm"` // 视频格式(mp4,flv)
	Duration string `json:"duration" db:"duration"` // 视频时长
}

// @Summary 上传视频信息
// @Description 普通用户可上传
// @Tags 视频
// @Accept json
// @Param Body body AddVideoParam true "desc"
// @Success 200 {string} json "{"message":"创建视频完成!","results":videoObject}"
// @Failure 400 {string} json "{"message": "参数错误/创建失败!", "results": null}"
// @Router /videos/ [post]
func AddVideo(ctx *gin.Context) {
	v := new(AddVideoParam)
	err := ctx.ShouldBindJSON(v)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	nv := new(models.Video)
	db.Move(v, nv, []string{})
	if ok := db.New(nv); ok {
		utils.HttpOk(ctx, "创建视频完成!", v)
	} else {
		utils.HttpBadRequest(ctx, "创建失败!", nil)
	}
}

// @Summary 删除视频
// @Description 只有管理员或作者可以删除视频
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Success 200 {string} json "{"message":"删除完成！","results": null}"
// @Failure 400 {string} json "{"message": "参数错误/删除失败!", "results": null}"
// @Router /videos/{id}/ [delete]
func DeleteVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	v := new(models.Video)
	db.Get(v, fmt.Sprintf("select * from video where id=%s", sid))
	if db.Delete("video", id) {
		// TODO 定期任务执行删除
		utils.HttpOk(ctx, "删除完成！", nil)
	}else{
		utils.HttpBadRequest(ctx, "删除失败！", nil)
	}
}

type ModifyVideoParam struct {
	Title string `json:"title" db:"title" binding:"required"` // 标题
	Desc string `json:"desc" db:"desc"` // 描述
	Category string `json:"category" db:"category"` // 视频类型
	Label string `json:"label" db:"label"` // 视频标签，多个自定义标签
	Rm string `json:"rm" db:"rm"` // 视频格式(mp4,flv)
	Duration string `json:"duration" db:"duration"` // 视频时长
}

// @Summary 修改视频
// @Description 只有管理员/或用户自己可以修改视频信息
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Param Body body ModifyVideoParam true "desc"
// @Success 200 {string} json "{"message":"成功修改用户!", "results": videoObject}"
// @Failure 400 {string} json "{"message": "参数错误", "results": null}"
// @Router /videos/{id}/ [put]
func ModifyVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	v := new(ModifyVideoParam)
	if err = ctx.ShouldBindJSON(v); err!=nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	GetSql := fmt.Sprintf("select * from video where id=%d", id)
	v2 := new(models.Video)
	if db.Get(v2, GetSql){
		db.Move(v, v2, []string{})
		if db.Modify(v2){
			utils.HttpOk(ctx, "成功修改视频信息!", v2)
		}else{
			utils.HttpBadRequest(ctx, "视频信息修改失败!", nil)
		}
	}else{
		utils.HttpBadRequest(ctx, "视频不存在!", nil)
	}
}

// @Summary 获取视频信息
// @Description
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Success 200 {string} json "{"message":"成功获取视频信息!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/视频不存在!", "results": null}"
// @Router /videos/{id}/ [get]
func GetVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	video := new(models.Video)
	GetSql := fmt.Sprintf("select * from video where id=%d", id)
	if db.Get(video, GetSql){
		utils.HttpOk(ctx, "成功获取视频信息!", video)
	}else{
		utils.HttpBadRequest(ctx, "视频不存在!", nil)
	}
}


// @Summary 获取视频信息列表
// @Description 无权限
// @Tags 视频
// @Accept json
// @Param page query string false "页数"
// @Param size query string false "个数"
// @Success 200 {string} json "{"message":"成功获取视频列表!","results":[{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}]}"
// @Failure 400 {string} json "{"msg": "参数错误/获取视频列表失败!", "results": null}"
// @Router /videos/ [get]
func GetVideos(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "0")
	sizeParam := ctx.DefaultQuery("size", "10")
	page, err1 := strconv.Atoi(pageParam)
	size, err2 := strconv.Atoi(sizeParam)
	if err1 != nil || err2 != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	selectSql := "select * from video"
	if page != 0{
		offset := (page-1) * size
		selectSql += fmt.Sprintf(" limit %d offset %d;", size, offset)
	}else{
		selectSql += ";"
	}
	var videos []models.Video
	if db.Select(&videos, selectSql){
		utils.HttpOk(ctx, "成功获取视频列表!", videos)
	}else{
		utils.HttpBadRequest(ctx, "获取视频列表失败!", nil)
	}
}


// @Summary 上传视频文件
// @Description 管理员或用户
// @Tags 视频
// @Accept mpfd
// @Param id path string true "视频id"
// @Param videoFile formData file false "视频文件"
// @Success 200 {string} json "{"message":"成功上传视频文件!文件地址:url!","results":null}"
// @Failure 400 {string} json "{"msg": "参数错误/视频不存在!/成功上传视频文件!但视频信息修改失败!", "results": null}"
// @Router /videos/{id}/upload/ [post]
func UploadVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	videoFile, err := ctx.FormFile("videoFile")
	if err != nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	v := new(models.Video)
	getSql := fmt.Sprintf("select * from video where id=%s", sid)
	if db.Get(v, getSql){

	}else{
		utils.HttpBadRequest(ctx, "参数错误!视频不存在！", nil)
	}
	// /videos/userid/sid_filename
	fileName := strings.Join([]string{sid, videoFile.Filename}, "_")
	filePath := path.Join(confs.Cfg["VIDEOS"], strconv.Itoa(v.UserId), fileName)
	err = ctx.SaveUploadedFile(videoFile, filePath)
	if err != nil {
		fmt.Println(filePath, err)
		utils.HttpBadRequest(ctx, "视频上传失败!", nil)
		return
	}
	v.VideoFile = strings.Split(filePath, "Helloc")[1]
	if db.Modify(v){
		utils.HttpOk(ctx, "成功上传视频!", v)
	}else{
		utils.HttpBadRequest(ctx, "视频上传成功，但信息更新失败!", nil)
	}
}

// @Summary 审核视频
// @Description 管理员
// @Tags 视频
// @Accept mpfd
// @Param id path string true "视频id"
// @Param pass formData bool true "是否通过"
// @Success 200 {string} json "{"message":"视频审核完成", "results":null}"
// @Failure 400 {string} json "{"msg": "视频审核失败!", "results": null}"
// @Router /videos/{id}/check/ [post]
func CheckVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	selectSql := fmt.Sprintf("select * from video where id=%d;", id)
	video := new(models.Video)
	if db.Get(video, selectSql){
		pass := ctx.DefaultPostForm("pass", "false")
		pass2 := false
		if pass == "true" {
			pass2 = true
		}
		video.Pass = pass2
		if db.Modify(video){
			utils.HttpOk(ctx, "视频审核完成!", nil)
		}else {
			utils.HttpBadRequest(ctx, "视频审核失败!", nil)
		}
		// TODO 审核不通过需要通知用户
	}else{
		utils.HttpBadRequest(ctx, "视频不存在!", nil)
	}
}
