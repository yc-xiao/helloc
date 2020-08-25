package confs

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path"
	"strings"
)

var Cfg = map[string]string{}

func init() {
	curPath, _ := os.Getwd()
	paths := strings.Split(curPath, "Helloc")
	baseFile := path.Join(paths[0], "Helloc/confs/base.ini")
	privateFile := path.Join(paths[0], "Helloc/confs/private.ini")
	fmt.Println(baseFile, privateFile)
	// 加载基础配置
	cfg, err := ini.Load(baseFile)
	if err != nil {
		log.Fatalf("Fail to parse 'confs/base.ini': %v", err)
	}
	for _, section := range cfg.Sections(){
		for _, key := range section.Keys(){
			Cfg[key.Name()] = key.Value()
		}
	}
	// 加载其他配置
	cfg, err = ini.Load(privateFile)
	if err != nil {
		log.Printf("Fail to parse 'confs/private.ini': %v", err)
	}else{
		for _, section := range cfg.Sections(){
			for _, key := range section.Keys(){
				Cfg[key.Name()] = key.Value()
			}
		}
	}

	// 加载文件配置
	rootPath := path.Join(paths[0], "Helloc")
	Cfg["ROOT"] = rootPath
	Cfg["STATIC"] = path.Join(rootPath, Cfg["STATIC"])
	Cfg["VIDEOS"] = path.Join(rootPath, Cfg["VIDEOS"])
	Cfg["IMAGES"] = path.Join(rootPath, Cfg["IMAGES"])
}

