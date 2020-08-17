package v1

import (
	. "Helloc/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUser(ctx *gin.Context){
	u := new(User)
	err := ctx.ShouldBindJSON(u)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if _, ok := UserCache[u.Id]; ok{
		HttpBadRequest(ctx, "用户已存在，创建失败!", nil)
	}else{
		UserCache[u.Id] = u
		HttpOk(ctx, "成功创建用户!", u)
	}
}

func DeleteUser(ctx *gin.Context){
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	delete(UserCache, id)
	HttpOk(ctx, "成功删除！", nil)
}

func ModifyUser(ctx *gin.Context){
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if _, ok := UserCache[id]; ok {
		u := new(User)
		ctx.ShouldBindJSON(u)
		UserCache[u.Id] = u
		HttpOk(ctx, "成功修改用户!", u)
	} else {
		HttpBadRequest(ctx, "用户不存在!", nil)
	}
}

func GetUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		HttpBadRequest(ctx, "用户不存在!", nil)
		return
	}
	if u, ok := UserCache[id]; ok {
		HttpOk(ctx, "成功获取用户!", u)
	} else {
		HttpBadRequest(ctx, "用户不存在!", nil)
	}
}

func GetUsers(ctx *gin.Context){
	pageParam := ctx.DefaultQuery("page", "0")
	sizeParam := ctx.DefaultQuery("size", "0")
	page, err1 := strconv.Atoi(pageParam)
	size, err2 := strconv.Atoi(sizeParam)
	if err1 != nil || err2 != nil {
		HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	if size == 0 || page == 0 {
		size = len(UserCache)
	}

	i := 0
	users := make([]*User, 0, size)
	for _, user := range UserCache {
		users = append(users, user)
		i++
		if i == size {
			break
		}
	}
	HttpOk(ctx, "成功获取用户列表!", users)
}

