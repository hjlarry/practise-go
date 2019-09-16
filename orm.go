package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/beego?charset=utf8")
	orm.RegisterModel(new(Author), new(User), new(Userinfo), new(Profile), new(Tag), new(Post))
	orm.Debug = true
	_ = orm.RunSyncdb("default", false, true)

}

func testOrm() {
	o := orm.NewOrm()
	//author := Author{Name: "hehe"}
	//
	//id, err := o.Insert(&author)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//author.Name = "ahah"
	//num, err := o.Update(&author)
	//fmt.Printf("Num: %d, ERR: %v\n", num, err)
	//
	//a := Author{Id: author.Id}
	//err = o.Read(&a)
	//fmt.Printf("author: %d, ERR: %v\n", a, err)

	// 插入数据
	var user User
	user.Name = "zxxx"
	id, err := o.Insert(&user)
	fmt.Println(id)
	fmt.Println(err)

	// 批量插入数据
	users := []User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(100, users)
	fmt.Println(successNums)
	fmt.Println(err)

	// 根据主键获取数据
	user = User{Id: 3}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}

	// 高级查询方式
	var user2 User
	err = o.QueryTable("user").Filter("Id", 1).One(&user2)
	fmt.Println(user2)

}

type Author struct {
	Id   int
	Name string `orm:"size(100)"`
}

type Userinfo struct {
	Uid        int `orm:"pk"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name       string
	Departname string
	Created    time.Time
}

type User struct {
	Id      int
	Name    string
	//Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	//User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}
