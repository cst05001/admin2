package models

import (
    "fmt"
    "log"
    "errors"
	"crypto/md5"
	"encoding/hex"
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func Connect() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
	orm.RegisterDataBase("default", "mysql", dsn)
}

//创建数据库
func createdb() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	var dsn string
	var sqlstring string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
    log.Println("dsn:", dsn)

	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()
}


func Syncdb() {
    createdb()
    Connect()

	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		log.Fatalln(err)
        return
	}

    createGroup("admin")
    g := GetGroupByGroupname("admin")
    createUser()
    u := GetUserByUsername("admin")
    GroupAddUser(g, &u)
    p := &Path{Pathname: "/admin"}
    AddPathWithGroup(p, g)
    p = &Path{Pathname: "/path/add"}
    AddPathWithGroup(p, g)
    p = &Path{Pathname: "/group/add"}
    AddPathWithGroup(p, g)
    p = &Path{Pathname: "/path/bindGroupAndPath"}
    AddPathWithGroup(p, g)
    createGroup("guest")
    g = GetGroupByGroupname("guest")
    p = &Path{Pathname: "/user/login"}
    AddPathWithGroup(p, g)
    p = &Path{Pathname: "/static"}
    AddPathWithGroup(p, g)
}

func createUser() {
    u := &User{}
    u.Username = "admin"
    u.Password = "admin"
    u.Nickname = "Mr. Fang"
    _, err := AddUser(u)
    if err != nil {
        log.Fatalln(err)
    }
}

func createGroup(groupname string) {
    g := &Group{}
    g.Groupname = groupname
    _, err := AddGroup(g)
    if err != nil {
        log.Fatalln(err)
    }
}

func createPath(path string) {
    p := &Path{}
    p.Pathname = path
    _, err := AddPath(p)
    if err != nil {
        log.Fatalln(err)
    }
}

func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func Init() {
    Connect()
}

func CheckLogin(username string, password string) (user User, err error) {
	user = GetUserByUsername(username)
	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != Strtomd5(password) {
		return user, errors.New("密码错误")
	}
	return user, nil
}
