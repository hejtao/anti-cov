package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Admin struct {
	Id int `json:"id"`

	Name string `json:"name"`

	AccountId  int       `json:"account_id"`
	Hidden     bool      `json:"hidden,omitempty"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`

	Messages []*Message `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Admin))
}

func CreateAdmin(admin *Admin) (int, error) {
	id, err := orm.NewOrm().Insert(admin)
	return int(id), err
}

func UpdateAdmin(admin *Admin, cols ...string) error {
	_, err := orm.NewOrm().Update(admin, cols...)
	return err
}

// 如果返回的 admin.Id == 0, 则出错
func GetAdminById(id int, rel ...interface{}) *Admin {
	o := orm.NewOrm()
	admin := new(Admin)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Admin)).Filter("id", id).One(admin)
	} else {
		_ = o.QueryTable(new(Admin)).Filter("id", id).RelatedSel(rel...).One(admin)
	}

	return admin
}

func GetAdminByAccountId(accountId int) *Admin {
	o := orm.NewOrm()
	admin := new(Admin)
	_ = o.QueryTable(new(Admin)).Filter("AccountId", accountId).RelatedSel().One(admin)

	return admin
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
	admins := make([]*Admin, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Admin)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&admins)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Admin)).SetCond(cond).Filter("hidden", false).Count()
		data["admins"] = admins

		return data
	}

	_, _ = o.QueryTable(new(Admin)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&admins)
	return admins
}
