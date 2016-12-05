package routers

import (
	"github.com/cst05001/admin2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{}, "get:LoginFront")
    beego.Router("/user/login", &controllers.UserController{}, "get:LoginFront")
    beego.Router("/user/login", &controllers.UserController{}, "post:Login")
    beego.Router("/admin", &controllers.AdminController{}, "get:Index")
}
