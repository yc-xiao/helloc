package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
}

func TestWrite(t *testing.T) {
	f, err := os.OpenFile("/home/youcan/Desktop/Go/works/Helloc/ab.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	fmt.Println(err)
	defer f.Close()
	fmt.Println("今天天气真好!")
	n, err := f.WriteString("今天天气真好!")
	fmt.Println(n, err)
}

func init(){
	fmt.Println("init1")
}
func init(){
	fmt.Println("init2")
}