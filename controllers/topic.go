package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] =CheckCount(this.Ctx)
	this.TplName="topic.html"
	this.Data["IsTopic"] = true
}
func (this *TopicController) Post()  {
	if !CheckCount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	title :=this.Input().Get("title")
	content :=this.Input().Get("content")
	err:=models.AddTopic(title,content)
	if err !=nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
	return
}
func (this *TopicController) Add()  {
	this.TplName="topic_add.html"
}