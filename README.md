#admin2

##用法

###角色

这个后台管理系统有 2 个角色，分别是 User、Group。 User 可以属于一个或者多个 Group。代码针对 Group 进行权限管控。主要管控语句有

* func RequireLogin(ctx *context.Context)
* func RequireGroup(ctx *context.Context, groupnames ...string)

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

***问：如何修改密码？***  
答：暂时没有，有没有熟悉前端 HTML、Bootstrap/VUE 的同学一起参与啊？？？？
