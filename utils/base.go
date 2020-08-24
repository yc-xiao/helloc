package utils

import (
	"Helloc/confs"
	"fmt"
	"os"
	"path"
)

func CreateUserStorageSpace(sid string) {
	videoPath := path.Join(confs.Cfg["VIDEOS"], sid)
	imagePath := path.Join(confs.Cfg["IMAGES"], sid)
	fmt.Println(videoPath, imagePath)
	if err := os.Mkdir(videoPath, os.ModePerm); err != nil{
		fmt.Println(err)
	}
	if err := os.Mkdir(imagePath, os.ModePerm); err != nil{
		fmt.Println(err)
	}
}

func ClearUserStorageSpace(sid string) {
	videoPath := path.Join(confs.Cfg["VIDEOS"], sid)
	imagePath := path.Join(confs.Cfg["IMAGES"], sid)
	fmt.Println(videoPath, imagePath)
	os.RemoveAll(videoPath)
	os.RemoveAll(imagePath)
}

