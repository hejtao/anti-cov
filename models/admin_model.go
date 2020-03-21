package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id int `json:"id"`

	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`

	Hidden     bool      `json:"hidden"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Admin))
}

func CreateAdmin(template *Admin) (int, error) {
	id, err := orm.NewOrm().Insert(template)
	return int(id), err
}

func UpdateAdmin(template *Admin, cols ...string) error {
	_, err := orm.NewOrm().Update(template, cols...)
	return err
}

// 如果返回的 template.Id == 0, 则出错
func GetAdminById(id int, rel ...interface{}) *Admin {
	o := orm.NewOrm()
	template := new(Admin)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Admin)).Filter("id", id).One(template)
	} else {
		_ = o.QueryTable(new(Admin)).Filter("id", id).RelatedSel(rel...).One(template)
	}

	return template
}

func DeleteAdminsByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Admin)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Admin)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetAdmins(cond *orm.Condition, page int, rel ...interface{}) interface{} {
	o := orm.NewOrm()
	templates := make([]*Admin, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Admin)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&templates)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Admin)).SetCond(cond).Filter("hidden", false).Count()
		data["templates"] = templates

		return data
	}

	_, _ = o.QueryTable(new(Admin)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&templates)
	return templates
}
