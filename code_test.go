package main

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
}

func init(){
	fmt.Println("init1")
}
func init(){
	fmt.Println("init2")
}