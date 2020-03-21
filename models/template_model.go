package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Template struct {
	Id int `json:"id"`

	Hidden     bool      `json:"hidden"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Template))
}

func CreateTemplate(template *Template) (int, error) {
	id, err := orm.NewOrm().Insert(template)
	return int(id), err
}

func UpdateTemplate(template *Template, cols ...string) error {
	_, err := orm.NewOrm().Update(template, cols...)
	return err
}

// 如果返回的 template.Id == 0, 则出错
func GetTemplateById(id int, rel ...interface{}) *Template {
	o := orm.NewOrm()
	template := new(Template)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Template)).Filter("id", id).One(template)
	} else {
		_ = o.QueryTable(new(Template)).Filter("id", id).RelatedSel(rel...).One(template)
	}

	return template
}

func DeleteTemplatesByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Template)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Template)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetTemplates(cond *orm.Condition, page int, rel ...interface{}) interface{} {
	o := orm.NewOrm()
	templates := make([]*Template, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Template)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&templates)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Template)).SetCond(cond).Filter("hidden", false).Count()
		data["templates"] = templates

		return data
	}

	_, _ = o.QueryTable(new(Template)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&templates)
	return templates
}

func TemplateExist(cond *orm.Condition) bool {
	return orm.NewOrm().QueryTable(new(Template)).SetCond(cond).Exist()
}
