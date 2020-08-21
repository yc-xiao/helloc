package utils

import (
	"Helloc/confs"
	"Helloc/models"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// 指定秘钥
var jwtSecret=[]byte(confs.Cfg["JWT_SECRET"])

type UserInfo struct {
	UserId int
	NickName string
	Account string
	IsAdmin bool
	jwt.StandardClaims
}

func DefaultGenerateJwt(user *models.User) (string, error){
	hourStr, _ := confs.Cfg["EXPIRES_HOURS"]
	hourInt, err := strconv.Atoi(hourStr)
	if err != nil {
		hourInt = 8
	}
	return GenerateJwt(user, hourInt)
}

func GenerateJwt(user *models.User, hourInt int)(string, error){
	// 1. 封装数据到jwt (过期时间，账号信息，)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserInfo{
		UserId: user.Id,
		Account: user.Account,
		NickName: user.NickName,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(hourInt) * time.Hour).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	})
	return token.SignedString(jwtSecret)
}

func CheckJwt(tokenString string) (*jwt.Token, *UserInfo, error){
	userInfo := new(UserInfo)
	// 第三个函数可以取消，即不加密
	token, err := jwt.ParseWithClaims(tokenString, userInfo, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return token, userInfo, err
}