package utils

import (
	"Helloc/models"
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	u := models.User{Name: "小明", Age: 18, Id: 1}
	token, err := GenerateJwt(&u)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
	token2, userInfo, err := CheckJwt(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token2, userInfo)
}
