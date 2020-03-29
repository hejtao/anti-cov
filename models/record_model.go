package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Record struct {
	Id int `json:"id"`

	Count int `json:"count"`
	Type  int `json:"type"`

	Hidden     bool      `json:"hidden"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

//func init() {
// orm.RegisterModel(new(Record))
//}

func CreateRecord(record *Record) (int, error) {
	id, err := orm.NewOrm().Insert(record)
	return int(id), err
}

func UpdateRecord(record *Record, cols ...string) error {
	_, err := orm.NewOrm().Update(record, cols...)
	return err
}

// 如果返回的 record.Id == 0, 则出错
func GetRecordById(id int, rel ...interface{}) *Record {
	o := orm.NewOrm()
	record := new(Record)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Record)).Filter("id", id).One(record)
	} else {
		_ = o.QueryTable(new(Record)).Filter("id", id).RelatedSel(rel...).One(record)
	}

	return record
}

func DeleteRecordsByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Record)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Record)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetRecords(cond *orm.Condition, page int, rel ...interface{}) interface{} {
	o := orm.NewOrm()
	records := make([]*Record, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Record)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&records)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Record)).SetCond(cond).Filter("hidden", false).Count()
		data["records"] = records

		return data
	}

	_, _ = o.QueryTable(new(Record)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&records)
	return records
}
