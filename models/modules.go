package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //下划线，只使用初始化函数。
	"os"
	"path"
	"strconv"
	"time"
	"fmt"
)
const (
	_DB_NAME="data/beego.db"//数据库名称
	_SQLITE3_DRIVER="sqlite3"//驱动名称
)

type Category struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCout int64
	TopiclsUserId int64
}
type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	views int64
	Author string
	ReplayTime time.Time `orm:"index"`
	ReplayCount int64
	ReplayLastUserId int64
}

func RegisterDB()  {
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Topic),new(Category))//注册模型
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)//注册驱动
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
	
}
func AddTopic(title,content string) error  {
	o:=orm.NewOrm()
	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
	}
	_,err:=o.Insert(topic)
	return err
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	words := ([]rune)(name)
	var inStr string
	for i := 0; i < len(words); i++ {
		fmt.Println(string(words[i : i+1]))
		inStr = inStr + string(words[i:i+1]) +","
	}

	cate := &Category{Title: name,Created:time.Now(),TopicTime:time.Now()}
	qs :=o.QueryTable("category")
	err:=qs.Filter("title",inStr).One(cate)
		if (err ==nil){//err ==nil 代表查到了。不需要添加因此返回。
		return err
	}
	_,err = o.Insert(cate)
	if(err !=nil){
		return err
	}
	return nil
}
func DelCategory(Id string) error {
	cid,err:=strconv.ParseInt(Id,10,64)
	if (err !=nil){
		return err
	}
	o:=orm.NewOrm()
	cate :=Category{Id:cid}
	o.Delete(&cate)
	return nil
}
func GetAllCategories() ([]Category,error) {
	o:=orm.NewOrm()
	qs:=o.QueryTable("category")
	Cate:=make([]Category,0)
	_,err:=qs.All(&Cate)
	return Cate,err
}