#admin2

##用法

###代码调用

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

###初始化数据库

在你的 beego 程序中执行参数

    ./YOUR_APP_NAME syncdb

默认账号密码都是 admin
