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
	tableName, fields, values := modelToItems(i)
	fieldString := strings.Join(fields[1:], ",")
	valueString := strings.Join(values[1:], ",")
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

func Modify(i interface{}) bool {
	tableName, fields, values := modelToItems(i)
	l := len(values)
	ss := make([]string, l, l)
	for i:=0; i<l; i++ {
		ss[i] = fmt.Sprintf("%s=%s", fields[i], values[i])
	}
	UpdateSql := fmt.Sprintf("update `%s` set %s where id=%s;", tableName, strings.Join(ss[1:], ","), values[0])
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

func modelToItems(i interface{}) (tableName string, fields []string, values []string){
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
	l := tElem.NumField()
	fields, values = make([]string, l), make([]string, l)
	for i:=0; i<l; i++ {
		tf, vf := tElem.Field(i), vElem.Field(i)
		fields[i] = tf.Tag.Get("db")
		// 临时处理
		switch tf.Type.String() {
		case "int":
			values[i] = strconv.Itoa(int(vf.Int()))
		case "bool":
			if vf.Bool(){
				values[i] = "1"
			}else{
				values[i] = "0"
			}
		default:
			values[i] = fmt.Sprintf(`"%s"`, vf.String())
		}
	}
	return
}