package main

import (
	"Helloc/models"
	"Helloc/models/utils"
)

func main() {
	//u := new(models.User)
	//t := reflect.TypeOf(u)
	//v := reflect.ValueOf(u)
	//tElem := t.Elem()
	//vElem := v.Elem()
	//ta, ok := utils.ModelToTable[vElem.Type().String()]
	//fmt.Println(ta, ok)
	//for i:=0; i<tElem.NumField(); i++ {
	//	tf, vf := tElem.Field(i), vElem.Field(i)
	//	fmt.Println(tf, vf, tf.Type)
	//
	//}
	u := models.User{
		NickName: "小明",
		Account: "小红",
		IsAdmin: true,
		CreatedTime: "",
	}
	utils.New(&u)
}
