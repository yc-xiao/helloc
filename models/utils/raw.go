package utils

import (
	"Helloc/confs"
	"Helloc/utils"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var DB *sqlx.DB

var ModelToTable = map[string]string {
	"models.User": "user",
	"models.Comment": "comment",
	"models.Video": "video",
}

func init() {
	dbSource, ok := confs.Cfg["DB_SOURCE"]
	if !ok {
		dbSource = "hello:123456@tcp(127.0.0.1:3306)/H"
	}
	db, err := sqlx.Open("mysql", dbSource)
	DB = db
	utils.ErrorBreak(err, "db连接失败")
}

func New(i interface{}) bool {
	tableName, AttrToValue := modelToItems(i, []string{})
	ll := len(AttrToValue)
	fields, values := make([]string, ll-1), make([]string, ll-1)

	j := 0
	for field, value := range AttrToValue {
		if field == "`id`" {
			continue
		}
		if field == "`createdTime`" {
			value = `"` + time.Now().Format("2006-01-02 15:04:05") + `"`
		}
		fields[j] = field
		values[j] = value
		j++
	}

	fieldString := strings.Join(fields, ",")
	valueString := strings.Join(values, ",")
	newSql := fmt.Sprintf("insert into `%s` (%s) values (%s);", tableName, fieldString, valueString)
	_, err := DB.Exec(newSql)
	log.Println(newSql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Delete(table string, id int) bool {
	deleteSql := fmt.Sprintf("delete from %s where id=%d", table, id)
	_, err := DB.Exec(deleteSql)
	log.Println(deleteSql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Modify(i interface{}, partFields []string) bool {
	// partFields 指定字段修改
	tableName, AttrToValue := modelToItems(i, partFields)
	fmt.Println(partFields, AttrToValue)
	ss := make([]string, len(AttrToValue)-1)
	id, _ := AttrToValue["`id`"]

	j := 0
	for attr, value := range AttrToValue {
		if attr == "`id`" {
			continue
		}
		ss[j] = fmt.Sprintf("%s=%s", attr, value)
		j++
	}

	UpdateSql := fmt.Sprintf("update `%s` set %s where id=%s;", tableName, strings.Join(ss, ","), id)
	_, err := DB.Exec(UpdateSql)
	log.Println(UpdateSql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Get(i interface{}, sql string) bool {
	log.Println(sql)
	if err := DB.Get(i, sql); err!=nil{
		log.Println(err)
		return false
	}
	return true
}

func Select(objs interface{}, sql string) bool{
	log.Println(sql)
	if err:=DB.Select(objs, sql); err!=nil{
		log.Println(err)
		return false
	}
	return true
}

func modelToItems(i interface{}, partFields []string) (tableName string, AttrToValue map[string]string){
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() != reflect.Ptr {
		utils.ErrorBreak(errors.New(" modelToItems错误，参数必须是指针类型!"), "")
	}
	tElem, vElem := t.Elem(), v.Elem()
	//fmt.Println(vElem.Type().String())
	tableName2, ok := ModelToTable[vElem.Type().String()]
	if !ok {
		utils.ErrorBreak(fmt.Errorf(" %s　表不存在!", vElem.Type()), "")
	}
	tableName = tableName2

	dic := make(map[string]bool)
	l2 := len(partFields)
	if l2 != 0 {
		for _, v := range partFields{
			dic[v] = true
		}
		if !dic["Id"] {
			dic["Id"] = true
		}
	}

	AttrToValue = make(map[string]string)
	var value string
	for i:=0; i<tElem.NumField(); i++ {
		tf, vf := tElem.Field(i), vElem.Field(i)
		if l2 != 0 && !dic[tf.Name] {
			continue
		}
		switch tf.Type.String() {
		case "int":
			value = strconv.Itoa(int(vf.Int()))
		case "bool":
			if vf.Bool(){
				value = "1"
			}else{
				value = "0"
			}
		default:
			value = fmt.Sprintf(`"%s"`, vf.String())
		}
		AttrToValue["`" + tf.Tag.Get("db") + "`"] = value
	}
	return
}

func Move(objParam interface{}, objModel interface{}, fields []string) []string{
	// 默认都是指针对象
	// 将 objParam　结构体的参数，取指定fields字段　赋予 objModel
	objPv := reflect.ValueOf(objParam)
	if objPv.Kind() == reflect.Ptr{
		objPv = objPv.Elem()
	}

	objMv := reflect.ValueOf(objModel).Elem()
	if len(fields) == 0{
		objPt := objPv.Type()
		for i:=0; i< objPt.NumField(); i++{
			f := objPt.Field(i)

			fields = append(fields, f.Name)
		}
	}
	//fmt.Println(fields)
	for _, field := range fields{
		//fmt.Println(field)
		v1 := objPv.FieldByName(field)
		v2 := objMv.FieldByName(field)

		if !v2.CanSet(){
			continue
		}
		v2.Set(v1)
	}
	//fmt.Println(objModel)
	return fields
}