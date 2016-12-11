# admin2

## 用法

    package controllers

    import (
        "github.com/astaxie/beego"
        "github.com/cst05001/admin2"
    )

    type MainController struct {
        beego.Controller
    }

    func (c *MainController) Get() {
        admin2.RequireGroup(c.Ctx, "admin")
        c.Data["Website"] = "beego.me"
        c.Data["Email"] = "astaxie@gmail.com"
        c.TplName = "index.tpl"
    }
