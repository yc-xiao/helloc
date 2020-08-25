package utils

import (
	"fmt"
	"testing"
)

func TestSendPhoneCode(t *testing.T) {
	SendPhoneCode("15390940293", "2333")
}

func TestGenerateRandNum(t *testing.T) {
	for i:=0 ; i< 10; i++ {
		fmt.Println(GenerateRandNum())
	}
}