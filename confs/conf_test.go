package confs

import (
	"fmt"
	"testing"
)

func TestCfg(t *testing.T) {
	for k,v := range Cfg{
		fmt.Println(k, v)
	}
}
