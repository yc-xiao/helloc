package v1

import (
	"Helloc/models"
	db "Helloc/models/utils"
	"Helloc/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)


type AddCommentParam struct {
	ObjId int `json:"objId" binding:"required" example:"1"` // 评论关联对象id
	Model string `json:"model" binding:"required" example:"user"` // 评论关联模型名
	Context string `json:"context" example:"context"` // 内容
	UserId int `json:"userId" binding:"required" example:"1"` // 创建用户id
	//OriginId int `json:"originId" binding:"required" example:"1"` // 最初评论id
	//OriginModel string `json:"originModel" binding:"required" example:"user"` // 最初模型
}

// @Summary 新增评论
// @Description 添加评论
// @Tags 评论
// @Accept json
// @Param Body body AddCommentParam true "desc"
// @Success 200 {object} utils.Response{results=models.Comment}
// @Failure 400 {object} utils.Response
// @Router /comments/ [post]
func AddComment(ctx *gin.Context) {
	c := new(AddCommentParam)
	err := ctx.ShouldBindJSON(c)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	nc := new(models.Comment)
	db.Move(c, nc, []string{})
	if ok := db.New(nc); ok {
		utils.HttpOk(ctx, "创建评论!", nc)
	} else {
		utils.HttpBadRequest(ctx, "评论失败!", nil)
	}
}

// @Summary 删除评论
// @Description 删除评论
// @Tags 评论
// @Accept json
// @Param id path string true "评论id"
// @Success 200 {string} json "{"results": null, "message":"删除完成"}"
// @Failure 400 {string} json "{"results": null, "message":"参数错误"}"
// @Router /comments/{id}/ [delete]
func DeleteComment(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	c := new(models.Comment)
	getSql := fmt.Sprintf("select * from comment where id=%d", id)
	if db.Get(c, getSql){

	}else{
		utils.HttpBadRequest(ctx, "评论不存在!", nil)
		return
	}
	if DeepDelete(c) {
		utils.HttpOk(ctx, "删除完成！", nil)
	}else{
		utils.HttpBadRequest(ctx, "删除失败！", nil)
	}

}

// @Summary 获取评论
// @Description desc
// @Tags 评论
// @Accept json
// @Param id path string true "评论id"
// @Success 200 {string} json "{"results": ["user", "comment", "video"]}, "message":""}"
// @Router /comments/{id}/ [get]
func GetComment(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}
	c := new(commentRes)
	GetSql := fmt.Sprintf("select * from comment where id=%d", id)
	if !db.Get(c, GetSql){
		utils.HttpBadRequest(ctx, "评论不存在!", nil)
		return
	}

	comments := []commentRes{*c}
	deepSelect(comments)
	utils.HttpOk(ctx, "成功获取评论!", comments)

}

type commentRes struct {
	models.Comment
	SecondComment []commentRes `json:"secondComment"` // 下级评论
}

func deepSelect(comments []commentRes) {
	for i:=0; i < len(comments); i++ {
		c := &comments[i]
		selectSql := fmt.Sprintf(`select * from comment where objId=%d and model="comment"`, c.Id)
		db.Select(&c.SecondComment, selectSql)
		deepSelect(c.SecondComment)
	}
}

func DeepDelete(c *models.Comment) bool {
	cursor, err := db.DB.Begin()
	if err != nil {
		log.Println("begin failed :", err)
		return false
	}
	if deepDelete(c, cursor) {
		cursor.Commit()
		return true
	}else{
		cursor.Rollback()
		return false
	}
}

func deepDelete(c *models.Comment, cursor *sql.Tx) bool {
	selectSql := fmt.Sprintf(`select * from comment where objId=%d and model="comment"`, c.Id)
	var cs []models.Comment
	if db.Select(&cs, selectSql){
		for _, c2 := range cs {
			if !deepDelete(&c2, cursor) {
				return false
			}
		}
	}else {
		return false
	}

	deleteSql := fmt.Sprintf("delete from comment where id=%d", c.Id)
	_, err := cursor.Exec(deleteSql)
	log.Println(deleteSql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// @Summary 获取视频评论
// @Description desc
// @Tags 评论
// @Accept json
// @Param vid path string true "视频id"
// @Param page query string false "页数"
// @Param size query string false "个数"
// @Success 200 {string} json "{"results": ["user", "comment", "video"]}, "message":""}"
// @Router /comment/video/{vid}/ [get]
func GetCommentByVideo(ctx *gin.Context) {
	// 通过视频id获取视频下的一级评论，在通过一级评论获取二级评论
	pageParam := ctx.DefaultQuery("page", "0")
	sizeParam := ctx.DefaultQuery("size", "10")
	page, err1 := strconv.Atoi(pageParam)
	size, err2 := strconv.Atoi(sizeParam)
	if err1 != nil || err2 != nil {
		utils.HttpBadRequest(ctx, "参数错误!", nil)
		return
	}

	vid := ctx.Param("vid")
	selectSql := fmt.Sprintf(`select * from comment where objId=%s and model="video"`, vid)
	if page != 0{
		offset := (page-1) * size
		selectSql += fmt.Sprintf(" limit %d offset %d;", size, offset)
	}else{
		selectSql += ";"
	}

	var comments []commentRes
	if db.Select(&comments, selectSql){
		// 递归查询很费时间，可以设置递归层数
		deepSelect(comments)
		utils.HttpOk(ctx, "成功获取评论!", comments)
	}else{
		utils.HttpBadRequest(ctx, "获取评论失败!", nil)
	}
}



// @Summary 获取评论关联模型表
// @Description desc
// @Tags 评论
// @Accept json
// @Success 200 {string} json "{"results": ["user", "comment", "video"]}, "message":""}"
// @Router /comment/models/ [get]
func GetModels(ctx *gin.Context) {
	utils.HttpOk(ctx, "获取评论可关联模型!", []string{"comment", "video"})
}

