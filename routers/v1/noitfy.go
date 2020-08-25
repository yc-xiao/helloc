package v1

import (
	db "Helloc/models/utils"
	"Helloc/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func SendCodeByPhone(ctx *gin.Context) {
	phone := ctx.Param("phone")
	code := utils.GenerateRandNum()
	if utils.SendPhoneCode(phone, code) {
		db.RSetExpireString(phone, code, 5*60)
		utils.HttpOk(ctx, "code:"+code, nil)
	} else {
		log.Println(phone, code)
		utils.HttpBadRequest(ctx, "发送失败", nil)
	}
}
