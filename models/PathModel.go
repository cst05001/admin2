package models

import (
    "log"
    "errors"
	"github.com/astaxie/beego/orm"
)

type Path struct {
    Id  int64
    Pathname    string `orm:"unique;size(128)" form:"pathname"`
    Group   []*Group `orm:"reverse(many)"`
}

func (this *Path) TableName() string {
    return "path"
}

func init() {
	orm.RegisterModel(new(Path))
}

func AddPath(p *Path) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func AddPathWithGroup(p *Path, g *Group) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	if err != nil {
        return id, err
    }
    p.Id = id
    err = o.Read(p, "Id")
	if err != nil {
        return id, err
    }
    m2m := o.QueryM2M(p, "Group")
    _, err = m2m.Add(g)
    return id, err
}

//更新用户
func UpdatePath(this *Path) (int64, error) {
	o := orm.NewOrm()
	path := make(orm.Params)
	if len(this.Pathname) > 0 {
		path["Pathname"] = this.Pathname
	}
	if len(path) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Path
	num, err := o.QueryTable(table).Filter("Id", this.Id).Update(path)
	return num, err
}

func DelPathById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Path{Id: Id})
	return status, err
}

func GetPathByPathname(pathname string) (path Path) {
	path = Path{Pathname: pathname}
	o := orm.NewOrm()
	o.Read(&path, "Pathname")
	return path
}

func GetAllPath() []*Path {
    paths := make([]*Path, 0)
    o := orm.NewOrm()
    qs := o.QueryTable("Path")
    _, err := qs.All(&paths)
    if err != nil {
        log.Println(err)
        return nil
    }
    for _, p := range(paths) {
        _, err = o.LoadRelated(p, "Group")
    }
    return paths
}

func GetPathById(id int64) (path Path) {
	path = Path{Id: id}
	o := orm.NewOrm()
	o.Read(&path, "Id")
	return path
}

func PathAddGroup(p *Path, g *Group) error {
    o := orm.NewOrm()
    m2m := o.QueryM2M(p, "Group")
    _, err := m2m.Add(g)
    return err
}

