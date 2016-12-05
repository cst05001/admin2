package models

import (
    "log"
    "errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type User struct {
    Id  int64
    Username    string `orm:"unique;size(32)" form:"username"  valid:"Required;MaxSize(20);MinSize(6)"`
    Password    string `orm:"size(32)" form:"password" valid:"Required;MaxSize(20);MinSize(6)"`
    Repassword  string `orm:"-" form:"repassword" valid:"Required"`
    Nickname    string `orm:"unique;size(32)" form:"nickname" valid:"Required;MaxSize(20);MinSize(2)"`
    Email       string `orm:"size(32)" form:"email" valid:"Email"`
    Group       []*Group `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}


//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email

	id, err := o.Insert(user)
	return id, err
}

//更新用户
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	err := o.QueryTable(&user).Filter("Username", username).One(&user)
    if err != nil {
        log.Println(err)
    }
    _, err = o.LoadRelated(&user, "Group")
    if err != nil {
        log.Println(err)
    }
	return user
}

func GetUserById(id int64) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	err := o.QueryTable(&user).Filter("Id", id).RelatedSel().One(&user)
    if err != nil {
        log.Println(err)
    }
    _, err = o.LoadRelated(&user, "Group")
    if err != nil {
        log.Println(err)
    }
	return user
}

func GroupAddUser(g *Group, u *User) error {
    o := orm.NewOrm()
    m2m := o.QueryM2M(u, "Group")
    _, err := m2m.Add(g)
    return err
}

