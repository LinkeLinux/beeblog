package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "index.tpl"
	t:=time.Now()
	s:=fmt.Sprintf("%d-%d-%d %d:%d:%d",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	c.Data["Date"]=s
	c.Data["IsLogin"]=CheckCount(c.Ctx)

}
