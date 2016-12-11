package admin2

import (
    "strings"
	"github.com/cst05001/admin2/controllers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
    "github.com/cst05001/admin2/models"
)

var pathList []*models.Path

func RequireLogin(ctx *context.Context) {
    user := getUser(ctx)
    if user == nil {
        ctx.Redirect(302, "/user/login")
        return
    }
    return
}

func RequireGroup(ctx *context.Context, groupname string) {
    user := getUser(ctx)
    if user == nil {
        ctx.Redirect(302, "/user/login")
        return
    }
    for _, user_group := range(user.Group) {
        if user_group.Groupname == groupname {
            return
        }
    }
    ctx.WriteString("权限不足")
    return
}

/*
func Check(ctx *context.Context) {
    is_white_list := isWhiteList(ctx)
    user := getUser(ctx)
    if user == nil && !is_white_list {
        ctx.Redirect(302, "/user/login")
        return
    } else if user != nil &&  !is_white_list {
        ctx.WriteString("权限不足")
        return
    }
    return
}
*/

func init() {
    models.Init()
    pathList = models.GetAllPath()
    registerRouter()
}

func registerRouter() {
    beego.Router("/user/login", &controllers.UserController{}, "get:LoginFront")
    beego.Router("/user/login", &controllers.UserController{}, "post:Login")
    beego.Router("/admin", &controllers.AdminController{}, "get:Index")
    beego.Router("/path/add", &controllers.PathController{}, "post:Add")
    beego.Router("/group/add", &controllers.GroupController{}, "post:Add")
    beego.Router("/path/bindGroupAndPath", &controllers.PathController{}, "post:AddGroupToPath")
    // beego.InsertFilter("/*", beego.BeforeRouter, Check)
}

func getPath(uri string) *models.Path {
    path := strings.Split(uri, "?")[0]
    for _, p := range(pathList) {
        if p.Pathname == path {
            return p
        }
    }
    return nil
}

func getUser(ctx *context.Context) *models.User {
    v := ctx.Input.Session("userinfo")
    if v != nil {
        return v.(*models.User)
    }
    return nil
}

/*
func isWhiteList(ctx *context.Context) bool {
    path := getPath(ctx.Request.RequestURI)
    if path == nil {
        return false
    }

    user := getUser(ctx)
    for _, group := range(path.Group) {
        if group.Groupname == "guest" {
            return true
        } else if user != nil {
            for _, user_group := range(user.Group) {
                if user_group.Groupname == group.Groupname {
                    return true
                }
            }
        }
    }
    return false
}
*/
