package confs

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var Cfg = map[string]string{}

func init() {
	var path string
	if path, _ = os.Getwd(); strings.Contains(path, "confs"){
		path = "base.ini"
	}else {
		path = "confs/base.ini"
	}

	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Fail to parse 'confs/base.ini': %v", err)
	}
	for _, section := range cfg.Sections(){
		for _, key := range section.Keys(){
			Cfg[key.Name()] = key.Value()
		}
	}
}
