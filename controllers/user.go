package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
    "github.com/cst05001/admin2/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) LoginFront() {
	this.TplName = "user_login.tpl"
}

func (this *UserController) Login() {
    u := models.User{}
    err := this.ParseForm(&u)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    u, err = models.CheckLogin(u.Username, u.Password)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    this.SetSession("userinfo", &u)
    this.Ctx.WriteString(fmt.Sprint("Welcome,", u.Nickname))
}

