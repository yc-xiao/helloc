package utils

import (
	"Helloc/models"
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	u := models.User{Id:1, NickName:"小明", Account: "xiaoming", IsAdmin: true}
	token, err := DefaultGenerateJwt(&u)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
	_, userInfo, err := CheckJwt(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(userInfo)
}
