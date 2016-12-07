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

func (this *PathController) AddGroupToPath() {
	pathname := this.GetString("pathname")
	groupname := this.GetString("groupname")

	path := models.GetPathByPathname(pathname)
	if path == nil {
		this.Ctx.WriteString("path 不存在")
		return
	}
	group := models.GetGroupByGroupname(groupname)
	if group == nil {
		this.Ctx.WriteString("group 不存在")
		return
	}

	err := models.PathAddGroup(path, group)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString("Succcessed")
}
