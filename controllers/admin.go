package controllers

import (
	"github.com/astaxie/beego"
    // "github.com/cst05001/admin2/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Index() {
    this.TplName = "admin_index.tpl"
}

