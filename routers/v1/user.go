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

type AddUserParam struct {
	Account string `json:"account" binding:"required" example:"xiaom"` // 账号
	NickName string `json:"nickname" binding:"required" example:"小明"` // 名称
	Password string `json:"password" binding:"required" example:"123456"` // 密码
	Email string `json:"email" example:" "` // 邮箱
	Phone string `json:"phone" example:" "` // 手机
	IsAdmin bool `json:"isAdmin" example:"false"` // 是否是管理员
	PhotoFile string `json:"photoFile" example:" "` // 用户头像
	CreatedTime string `json:"createdTime" example:" "` // 创建时间
}

// @Summary 新增用户
// @Description 只有管理员可以添加用户
// @Tags 用户
// @Accept json
// @Param Body body AddUserParam true "desc"
// @Success 200 {object} utils.Response{results=models.User}
// @Failure 400 {object} utils.Response
// @Router /users/ [post]
func AddUser(ctx *gin.Context) {
	u := new(AddUserParam)
	err := ctx.ShouldBindJSON(u)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	nu := new(models.User)
	db.Move(u, nu, []string{})
	if ok := db.New(nu); ok {
		utils.CreateUserStorageSpace(strconv.Itoa(nu.Id))
		utils.HttpOk(ctx, "创建用户完成!", u)
	} else {
		utils.HttpBadRequest(ctx, "创建失败，用户已存在!", nil)
	}
}

// @Summary 删除用户
// @Description 只有管理员可以删除用户
// @Tags 用户
// @Accept json
// @Param id path string true "用户id"
// @Success 200 {string} json "{"message":"删除完成！","results": null}"
// @Failure 400 {string} json "{"message": "参数错误/删除失败!", "results": null}"
// @Router /users/{id}/ [delete]
func DeleteUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if db.Delete("user", id) {
		utils.ClearUserStorageSpace(sid)
		utils.HttpOk(ctx, "删除完成！", nil)
	}else{
		utils.HttpBadRequest(ctx, "删除失败！", nil)
	}
}

type ModifyUserParam struct {
	NickName string `json:"nickname" binding:"required" example:"xiaom"` // 名称
	Password string `json:"password" binding:"required" example:"123456"` // 密码
	Email string `json:"email" example:" "` // 邮箱
	Phone string `json:"phone" example:" "` // 手机
	PhotoFile string `json:"phone" example:"/static/images/1/1.png"` // 头像
}
// @Summary 修改用户
// @Description 只有管理员/或用户自己可以修改
// @Tags 用户
// @Accept json
// @Param id path string true "用户id" default(1)
// @Param body body ModifyUserParam true "desc"
// @Success 200 {string} json "{"message":"成功修改用户!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/用户不存在!/用户修改失败!", "results": null}"
// @Router /users/{id}/ [put]
func ModifyUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	u := new(ModifyUserParam)
	if err = ctx.ShouldBindJSON(u); err!=nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	user := new(models.User)
	if db.Get(user, GetSql){
		db.Move(u, user, []string{})
		if db.Modify(user){
			utils.HttpOk(ctx, "成功修改用户!", user)
		}else{
			utils.HttpBadRequest(ctx, "用户修改失败!", nil)
		}
	}else{
		utils.HttpBadRequest(ctx, "用户不存在!", nil)
	}
}

// @Summary 获取用户
// @Description 登陆后才允许获取用户信息
// @Tags 用户
// @Accept json
// @Param id path string true "用户id"
// @Success 200 {string} json "{"message":"成功获取用户!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/用户不存在!", "results": null}"
// @Router /users/{id}/ [get]
func GetUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	u := new(models.User)
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	if db.Get(u, GetSql){
		utils.HttpOk(ctx, "成功获取用户!", u)
	}else{
		utils.HttpBadRequest(ctx, "用户不存在!", nil)
	}
}


// @Summary 获取用户列表
// @Description 暂时无权限/后需要登录
// @Tags 用户
// @Accept json
// @Param page query string false "页数"
// @Param size query string false "个数"
// @Success 200 {string} json "{"message":"成功获取用户列表!","results":[{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}]}"
// @Failure 400 {string} json "{"msg": "参数错误/获取用户列表失败!", "results": null}"
// @Router /users/ [get]
func GetUsers(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "0")
	sizeParam := ctx.DefaultQuery("size", "10")
	page, err1 := strconv.Atoi(pageParam)
	size, err2 := strconv.Atoi(sizeParam)
	if err1 != nil || err2 != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	selectSql := "select * from user"
	if page != 0{
		offset := (page-1) * size
		selectSql += fmt.Sprintf(" limit %d offset %d;", size, offset)
	}else{
		selectSql += ";"
	}
	var users []models.User
	if db.Select(&users, selectSql){
		utils.HttpOk(ctx, "成功获取用户列表!", users)
	}else{
		utils.HttpBadRequest(ctx, "获取用户列表失败!", nil)
	}
}


// @Summary 上传用户头像
// @Description 只有管理员/或用户自己可以修改
// @Tags 用户
// @Accept mpfd
// @Param id path string true "用户id"
// @Param photoFile formData file true "头像"
// @Success 200 {string} json "{"message":"成功上传用户头像!","results":[{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}]}"
// @Failure 400 {string} json "{"msg": "参数错误/头像上传失败!", "results": null}"
// @Router /users/{id}/photo/ [put]
func UploadPhoto(ctx *gin.Context) {
	sid := ctx.Param("id")
	photoFile, err := ctx.FormFile("photoFile")
	if err != nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	u := new(models.User)
	getSql := fmt.Sprintf("select * from user where id=%s", sid)
	if db.Get(u, getSql){

	}else{
		utils.HttpBadRequest(ctx, "参数错误!用户不存在！", nil)
	}

	images := confs.Cfg["IMAGES"]
	filePath := path.Join(images, sid, photoFile.Filename)
	err = ctx.SaveUploadedFile(photoFile, filePath)
	if err != nil {
		fmt.Println(filePath, err)
		utils.HttpBadRequest(ctx, "头像上传失败!", nil)
		return
	}
	u.Phone = strings.Split(filePath, "Helloc")[1]
	if db.Modify(u){
		utils.HttpOk(ctx, "成功上传用户头像!", u)
	}else{
		utils.HttpBadRequest(ctx, "头像上传成功，但信息更新失败!", nil)
	}
}
