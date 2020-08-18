package main

import (
	"Helloc/models"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for i:=0 ; i <= 20; i++ {
		u, _ := models.UserCache[i]
		fmt.Println(u.IsAdmin)
	}
}
