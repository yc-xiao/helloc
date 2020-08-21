package utils

import "log"

func ErrorBreak(err error, msg string) {
	if err != nil {
		log.Fatal(msg, " -> ", err)
	}
}

func FalseBreak(b bool, msg string){
	if !b {
		log.Println(msg)
	}
}
