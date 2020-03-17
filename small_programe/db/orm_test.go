package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/beego?charset=utf8")
	orm.RegisterModel(new(Author), new(User), new(Userinfo), new(Profile), new(Tag), new(Post))
	orm.Debug = true
	_ = orm.RunSyncdb("default", false, true)

}

func TestOrm(t *testing.T) {
	o := orm.NewOrm()
	author := Author{Name: "hehe"}

	id, err := o.Insert(&author)
	t.Logf("ID: %d, ERR: %v\n", id, err)

	author.Name = "ahah"
	num, err := o.Update(&author)
	t.Logf("Num: %d, ERR: %v\n", num, err)

	a := Author{Id: author.Id}
	err = o.Read(&a)
	t.Logf("author: %v, ERR: %v\n", a, err)

	// 插入数据
	var user User
	user.Name = "zxxx"
	id, err = o.Insert(&user)
	t.Log(id)
	t.Log(err)

	// 批量插入数据
	users := []User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(100, users)
	t.Log(successNums)
	t.Log(err)

	// 根据主键获取数据
	user = User{Id: 3}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		t.Log("查询不到")
	} else if err == orm.ErrMissPK {
		t.Log("找不到主键")
	} else {
		t.Log(user.Id, user.Name)
	}

	// 高级查询方式
	var user2 User
	err = o.QueryTable("user").Filter("Id", 1).One(&user2)
	t.Log(user2)

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
	Id   int
	Name string
	//Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id  int
	Age int16
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
