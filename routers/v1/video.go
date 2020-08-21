package v1

import (
	"Helloc/confs"
	"Helloc/models"
	db "Helloc/models/utils"
	. "Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"strings"
)


// @Summary 上传视频信息
// @Description 普通用户可上传
// @Tags 视频
// @Accept json
// @Param body body string true "title(*标题)" default({"title": "小明大战小黄", "desc": "test", "category": "电影", "label": "小明，小黄，动作，武打", "rm": "mp4"})
// @Success 200 {string} json "{"message":"创建视频完成!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/创建失败!", "results": null}"
// @Router /videos/ [post]
func AddVideo(ctx *gin.Context) {
	v := new(models.Video)
	err := ctx.ShouldBindJSON(v)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if ok := db.New(v); ok {
		HttpOk(ctx, "创建视频完成!", v)
	} else {
		HttpBadRequest(ctx, "创建失败!", nil)
	}
}

// @Summary 删除视频
// @Description 只有管理员或作者可以删除视频
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Success 200 {string} json "{"message":"删除完成！","results": null}"
// @Failure 400 {string} json "{"msg": "参数错误/删除失败!", "results": null}"
// @Router /videos/{id}/ [delete]
func DeleteVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if db.Delete("video", id) {
		// TODO 删除视频原文件
		HttpOk(ctx, "删除完成！", nil)
	}else{
		HttpBadRequest(ctx, "删除失败！", nil)
	}
}

// @Summary 修改视频
// @Description 只有管理员/或用户自己可以修改视频信息
// @Tags 视频
// @Accept json
// @Param body body string true "title(*标题)" default({"title": "小明大战小黄", "desc": "test", "category": "电影", "label": "小明，小黄，动作，武打", "rm": "mp4"})
// @Param body body string true "用户信息" default({"nickname": "小明", "account": "xiaoming", "password": "123456", "email": "", "phone": ""})
// @Success 200 {string} json "{"message":"成功修改用户!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/用户不存在!/用户修改失败!", "results": null}"
// @Router /videos/{id}/ [put]
func ModifyVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	video := new(models.Video)
	if err = ctx.ShouldBindJSON(video); err!=nil{
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	GetSql := fmt.Sprintf("select * from video where id=%d", id)
	video2 := new(models.Video)
	if db.Get(video2, GetSql){
		video2.Title, video2.Desc, video2.Category, video2.Label, video2.Rm = video.Title, video.Desc, video.Category, video2.Label, video2.Rm
		if db.Modify(video2){
			HttpOk(ctx, "成功修改视频信息!", video2)
		}else{
			HttpBadRequest(ctx, "视频信息修改失败!", nil)
		}
	}else{
		HttpBadRequest(ctx, "视频不存在!", nil)
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
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	video := new(models.Video)
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	if db.Get(video, GetSql){
		HttpOk(ctx, "成功获取视频信息!", video)
	}else{
		HttpBadRequest(ctx, "视频不存在!", nil)
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
		HttpBadRequest(ctx, "参数错误!", nil)
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
		HttpOk(ctx, "成功获取视频列表!", videos)
	}else{
		HttpBadRequest(ctx, "获取视频列表失败!", nil)
	}
}


// @Summary 上传视频文件
// @Description 管理员或用户
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Param videoFile formData file false "视频文件"
// @Success 200 {string} json "{"message":"成功上传视频文件!文件地址:url!","results":null}"
// @Failure 400 {string} json "{"msg": "参数错误/视频不存在!/成功上传视频文件!但视频信息修改失败!", "results": null}"
// @Router /{id}/upload/ [post]
func UploadVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil{
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	selectSql := fmt.Sprintf("select * from video where id=%d;", id)
	video := new(models.Video)
	if db.Get(video, selectSql){
		videoFile := confs.Cfg["VIDEOS"]
		_, fileHeader, _ := ctx.Request.FormFile("videoFile")
		videoFile = path.Join(videoFile, strconv.Itoa(video.UserId), fileHeader.Filename)
		err := ctx.SaveUploadedFile(fileHeader, videoFile)
		if err != nil {
			HttpBadRequest(ctx, "成功上传视频文件!但视频信息修改失败!", nil)
			return
		}
		videoPaths := strings.Split(videoFile, "Helloc")
		video.VideoFile = videoPaths[1]
		if db.Modify(video) {
			HttpOk(ctx, "成功上传视频文件!文件地址:" + videoFile, nil)
		}else{
			HttpBadRequest(ctx, "成功上传视频文件!但视频信息修改失败!", nil)
		}
	}else{
		HttpBadRequest(ctx, "视频不存在!", nil)
	}
}

// @Summary 审核视频
// @Description 管理员
// @Tags 视频
// @Accept json
// @Param id path string true "视频id"
// @Param pass formData bool true "是否通过"
// @Success 200 {string} json "{"message":"视频审核完成", "results":null}"
// @Failure 400 {string} json "{"msg": "视频审核失败!", "results": null}"
// @Router /{id}/upload/ [post]
func CheckVideo(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil{
		HttpBadRequest(ctx, "参数错误!", nil)
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
			HttpOk(ctx, "视频审核完成!", nil)
		}else {
			HttpBadRequest(ctx, "视频审核失败!", nil)
		}
		// TODO 审核不通过需要通知用户
	}else{
		HttpBadRequest(ctx, "视频不存在!", nil)
	}
}
