package models

import (
    "errors"
	"github.com/astaxie/beego/orm"
)

type Group struct {
    Id  int64
    Groupname    string `orm:"unique;size(32)" form:"groupname"  valid:"Required`
    User    []*User `orm:"reverse(many)"`
    Path    []*Path `orm:"rel(m2m)"`
}

func (this *Group) TableName() string {
    return "group"
}

func init() {
	orm.RegisterModel(new(Group))
}

func AddGroup(g *Group) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(g)
	return id, err
}

//更新用户
func UpdateGroup(this *Group) (int64, error) {
	o := orm.NewOrm()
	group := make(orm.Params)
	if len(this.Groupname) > 0 {
		group["Groupname"] = this.Groupname
	}
	if len(group) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Group
	num, err := o.QueryTable(table).Filter("Id", this.Id).Update(group)
	return num, err
}

func DelGroupById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Group{Id: Id})
	return status, err
}

func GetGroupByGroupname(groupname string) (group *Group) {
	group = &Group{Groupname: groupname}
	o := orm.NewOrm()
	err := o.Read(group, "Groupname")
    if err != nil {
        return nil
    }
	return group
}

func GetGroupById(id int64) (group Group) {
	group = Group{Id: id}
	o := orm.NewOrm()
	o.Read(&group, "Id")
	return group
}
