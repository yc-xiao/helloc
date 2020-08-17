package utils

import "Helloc/confs"

// 指定秘钥
var jwtSecret=[]byte(confs.Cfg["JWT_SECRET"])

type Account struct {
	Account string `json:"account"`
	Password string `json:"password"`
}