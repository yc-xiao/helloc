package utils

import (
	"Helloc/confs"
	"Helloc/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 指定秘钥
var jwtSecret=[]byte(confs.Cfg["JWT_SECRET"])

type UserInfo struct {
	UserId int
	UserName string
	jwt.StandardClaims
}

func generateJwt(user *models.User) (string, error){
	// 1. 封装数据到jwt (过期时间，账号信息，)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserInfo{
		UserId: user.Id,
		UserName: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	})
	return token.SignedString(jwtSecret)
}

func checkJwt(tokenString string) (*jwt.Token, *UserInfo, error){
	userInfo := new(UserInfo)
	// 第三个函数可以取消，即不加密
	token, err := jwt.ParseWithClaims(tokenString, userInfo, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return token, userInfo, err
}