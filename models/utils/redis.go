package utils

import (
	"Helloc/utils"
	"github.com/garyburd/redigo/redis"
	"log"
)

var obj redis.Conn

func init() {
	c, err := redis.Dial("tcp","localhost:6379")
	if err != nil{
		utils.ErrorBreak(err, "redis 连接失败!")
	}
	obj = c
}

func RSetString(key, value string) bool {
	_, err := obj.Do("Set", key, value)
	if err != nil {
		log.Println(err, "redis 存储错误！")
		return false
	}
	return true
}

func RGetString(key string) (string, error) {
	return redis.String(obj.Do("Get", key))
}

func RSetExpireString(key, value string, expireSecond int) bool {
	if !RSetString(key, value){
		return false
	}
	_, err := obj.Do("expire", key, expireSecond)
	if err != nil {
		log.Println(err, "redis 设置过期时间错误！")
		return false
	}
	return true
}

func RedisClose() {
	obj.Close()
}