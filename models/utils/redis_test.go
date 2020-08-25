package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestRGetString(t *testing.T) {
	s, e :=  RGetString("hello3")
	fmt.Println(s, e)
}

func TestRSetString(t *testing.T) {
	if RSetString("hello", "233333"){
		fmt.Println(RGetString("hello"))
	}else {
		fmt.Println("error")
	}
}

func TestRSetExpireString(t *testing.T) {
	if RSetExpireString("hello2", "helloc2", 1){
		fmt.Println(RGetString("hello2"))
		time.Sleep(2*time.Second)
		fmt.Println(RGetString("hello2"))
	}else {
		fmt.Println("error")
	}
}