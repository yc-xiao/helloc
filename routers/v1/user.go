package v1

import (
	"Helloc/models"
	db "Helloc/models/utils"
	"Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddUserParam struct {
	Account string `json:"account" binding:"required" example:"小明"`
	Nickname string `json:"nickname" binding:"required" example:"xiaom"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email string `json:"email" example:""`
	Phone string `json:"phone" example:""`
	IsAdmin bool `json:"isAdmin" example:"false"`
}


// @Summary 新增用户
// @Description 只有管理员可以添加用户
// @Tags 用户
// @Accept json
// @Param Body body AddUserParam true "desc"
// @Success 200 {object} utils.ResponseStruct
// @Failure 400 {object} utils.ResponseStruct
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
	if ok := db.New(u); ok {
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
// @Failure 400 {string} json "{"msg": "参数错误/删除失败!", "results": null}"
// @Router /users/{id}/ [delete]
func DeleteUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if db.Delete("user", id) {
		utils.HttpOk(ctx, "删除完成！", nil)
	}else{
		utils.HttpBadRequest(ctx, "删除失败！", nil)
	}
}

// @Summary 修改用户
// @Description 只有管理员/或用户自己可以修改
// @Tags 用户
// @Accept json
// @Param id path string true "用户id" default(1)
// @Param body body string true "用户信息" default({"nickname": "小明", "password": "123456", "email": "", "phone": ""})
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
	u := new(models.User)
	if err = ctx.ShouldBindJSON(u); err!=nil{
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	user := new(models.User)
	if db.Get(user, GetSql){
		// 只允许修改NickName, Password, Email, Phone
		user.NickName, user.Password, user.Email, user.Phone = u.NickName, u.Password, u.Email, u.Phone
		if db.Modify(user){
			utils.HttpOk(ctx, "成功修改用户!", u)
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
