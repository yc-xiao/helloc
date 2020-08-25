package utils

import (
	"Helloc/models"
	"fmt"
	"testing"
)

func TestRExce(t *testing.T) {
	insertSql := fmt.Sprintf(`delete from user where account in ("x", "x2");`)
	r, err := DB.Exec(insertSql)
	fmt.Println(r.LastInsertId())
	fmt.Println(r.RowsAffected())
	fmt.Println(r, err)
}

func TestRQuery(t *testing.T) {
	querySql := fmt.Sprintf(`select id from user;`)
	q, err := DB.Query(querySql)
	fmt.Println(q, err)

	defer q.Close() // 需要释放连接
	for q.Next(){
		u := new(models.User)
		err = q.Scan(&u.Id)
		fmt.Println(err)
		fmt.Println(u.Id, u)
	}
}


func TestRSelect(t *testing.T) {
	us := []models.User{}
	authSql := fmt.Sprintf(`select * from user where account="%s" and password="%s";`, "xiaoming", "123456")
	fmt.Println(authSql)
	err := DB.Select(&us, authSql)
	fmt.Println(err)
	fmt.Println(us)
}

func TestRGet(t *testing.T) {
	u := new(models.User)
	type user struct {
		account string
		password string
	}
	ss := []struct{
		account string
		password string
	}{{"xiaohong", "123456"}, {"xiaoming", "1234567"}}

	for _, s := range ss {
		authSql := fmt.Sprintf(`select * from user where account="%s" and password="%s";`, s.account, s.password)
		err := DB.Get(u, authSql)
		if err != nil {
			fmt.Println(authSql)
			fmt.Println(err)
		}else {
			fmt.Println(u)
		}
		fmt.Println()
	}
}

func TestNew(t *testing.T) {
	funcName := "New"
	u := models.User{
		NickName: "小明",
		Account: "小红",
		IsAdmin: true,
		CreatedTime: "",
	}
	if New(&u) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}
}

func TestDelete(t *testing.T) {
	funcName := "Delete"
	if Delete("user", 9) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}
}

func TestModify(t *testing.T) {
	funcName := "New"
	u := models.User{
		NickName: "小明2",
		Account: "xiaoming2",
		IsAdmin: false,
		Password: "1234567",
		CreatedTime: "",
		Id: 1,
	}
	if Modify(&u, []string{"NickName", "Account", "Password"}) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}

	if Modify(&u, []string{}) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}
}

func TestGet(t *testing.T) {
	funcName := "Get"
	sql := "select * from user where id=1;"
	u := new(models.User)
	if Get(u, sql) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}
	fmt.Println(u)
}

func TestSelect(t *testing.T) {
	funcName := "Select"
	var users []models.User
	sql := "select * from user;"
	if Select(&users, sql) {
		t.Logf("func %s is ok, Nice!!!!!!", funcName)
	}else {
		t.Errorf("func %s is fail, NO!!!!", funcName)
	}
	fmt.Println(users)
}

func TestMove(t *testing.T) {
	u1 := new(models.User)
	u1.Id = 2
	u1.IsAdmin = true
	u1.Phone = "2333"
	u1.NickName = "小米"
	u2 := new(models.User)
	fmt.Println(u1, u2)
	ss := []string{"Id", "NickName", "IsAdmin"}
	Move(u1, u2, ss)
	fmt.Println(u2)
	u3 := models.User{Id: 3}
	Move(u3, u2, ss)
	fmt.Println(u2)
}