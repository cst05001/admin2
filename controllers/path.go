package controllers

import (
	"github.com/astaxie/beego"
    "github.com/cst05001/admin2/models"
)

type PathController struct {
	beego.Controller
}

func (this *PathController) AddFront() {
	this.TplName = "path_add.tpl"
}

func (this *PathController) Add() {
    p := &models.Path{}
    err := this.ParseForm(p)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    models.AddPath(p)
    this.Ctx.WriteString("OK")
}

