package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	cache := map[int]string{1:"1", 2:"2"}
	a := [10]string{}
	i := 0
	for _, v := range cache{
		a[i] = v
		i++
		if i == 1 {
			break
		}
	}
	fmt.Println(a)
}
