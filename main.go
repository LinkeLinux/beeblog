package main

import (
	"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	models.RegisterDB()
}
func main() {

	orm.RunSyncdb("default",false,true)//检查表是否重建，第二个参数false，只新建一次，如果为True则如果存在，删除之后新建。
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/topic",&controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Run()
}

