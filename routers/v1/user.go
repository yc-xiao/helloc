package v1

import (
	"Helloc/models"
	db "Helloc/models/utils"
	. "Helloc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 新增用户
// @Description 只有管理员可以添加用户
// @Tags 用户
// @Accept json
// @Param body body string true "account(*账号), passowrd(*密码) " default({"nickname": "小明", "account": "xiaom", "password": "123456", "email": "", "phone": "", "isAdmin": false, "photoFile":""})
// @Success 200 {string} json "{"message":"创建用户完成!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/用户已存在，创建失败!", "results": null}"
// @Router /users/ [post]
func AddUser(ctx *gin.Context) {
	u := new(models.User)
	err := ctx.ShouldBindJSON(u)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if ok := db.New(u); ok {
		HttpOk(ctx, "创建用户完成!", u)

	} else {
		HttpBadRequest(ctx, "用户已存在，创建失败!", nil)
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
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if db.Delete("user", id) {
		HttpOk(ctx, "删除完成！", nil)
	}else{
		HttpBadRequest(ctx, "删除失败！", nil)
	}
}

// @Summary 修改用户
// @Description 只有管理员/或用户自己可以修改
// @Tags 用户
// @Accept json
// @Param id path string true "用户id" default(1)
// @Param body body string true "用户信息" default({"nickname": "小明", "account": "xiaoming", "password": "123456", "email": "", "phone": ""})
// @Success 200 {string} json "{"message":"成功修改用户!","results":{"id":1,"nickname":"小明","account":"xiaoming","password":"123456","email":"","phone":"","isAdmin":false,"photoFile":"","createdTime":""}}}"
// @Failure 400 {string} json "{"msg": "参数错误/用户不存在!/用户修改失败!", "results": null}"
// @Router /users/{id}/ [put]
func ModifyUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	u := new(models.User)
	u.Id = id
	if err = ctx.ShouldBindJSON(u); err!=nil{
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	if db.Get(new(models.User), GetSql){
		if db.Modify(u){
			HttpOk(ctx, "成功修改用户!", u)
		}else{
			HttpBadRequest(ctx, "用户修改失败!", nil)
		}
	}else{
		HttpBadRequest(ctx, "用户不存在!", nil)
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
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	u := new(models.User)
	GetSql := fmt.Sprintf("select * from user where id=%d", id)
	if db.Get(u, GetSql){
		HttpOk(ctx, "成功获取用户!", u)
	}else{
		HttpBadRequest(ctx, "用户不存在!", nil)
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
		HttpBadRequest(ctx, "参数错误!", nil)
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
		HttpOk(ctx, "成功获取用户列表!", users)
	}else{
		HttpBadRequest(ctx, "获取用户列表失败!", nil)
	}
}
