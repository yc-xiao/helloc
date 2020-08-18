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
	filePath := path.Join(paths[0], "Helloc/confs/base.ini")
	fmt.Println(filePath)

	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Fatalf("Fail to parse 'confs/base.ini': %v", err)
	}
	for _, section := range cfg.Sections(){
		for _, key := range section.Keys(){
			Cfg[key.Name()] = key.Value()
		}
	}
}
