package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestSQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET name=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	t.Log(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set name=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	t.Log(affect)

	//	查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		t.Log(uid)
		t.Log(username)
		t.Log(department)
		t.Log(created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
