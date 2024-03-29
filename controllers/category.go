package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"

)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get()  {
	op := this.Input().Get("op")
	switch op {
	case "add":
		name :=this.Input().Get("name")
		if (len(name) ==0){
			break
		}

		err :=models.AddCategory(name)
		if(err !=nil){
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	case "del":
		id :=this.Input().Get("id")
		err :=models.DelCategory(id)
		if (err !=nil){
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	}
	this.TplName="category.html"
	this.Data["IsCategory"]= true

	var err error
	this.Data["Categories"],err=models.GetAllCategories()
	if(err !=nil) {
		beego.Error(err)
	}
}
