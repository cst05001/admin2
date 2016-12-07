package controllers

import (
	"github.com/astaxie/beego"
    "github.com/cst05001/admin2/models"
)

type GroupController struct {
	beego.Controller
}

func (this *GroupController) AddFront() {
	this.TplName = "path_add.tpl"
}

func (this *GroupController) Add() {
    group := &models.Group{}
    err := this.ParseForm(group)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    _, err = models.AddGroup(group)
	if err == nil {
    	this.Ctx.WriteString("OK")
	} else {
		this.Ctx.WriteString(err.Error())
	}
}
