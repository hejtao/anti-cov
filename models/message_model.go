package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id int `json:"id"`

	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content" orm:"size(2047)"`

	Hidden     bool      `json:"hidden,omitempty"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`

	Admin *Admin `orm:"rel(fk);null"`
}

func init() {
	orm.RegisterModel(new(Message))
}

func CreateMessage(message *Message) (int, error) {
	id, err := orm.NewOrm().Insert(message)
	return int(id), err
}

func UpdateMessage(message *Message, cols ...string) error {
	_, err := orm.NewOrm().Update(message, cols...)
	return err
}

// 如果返回的 message.Id == 0, 则出错
func GetMessageById(id int, rel ...interface{}) *Message {
	o := orm.NewOrm()
	message := new(Message)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Message)).Filter("id", id).One(message)
	} else {
		_ = o.QueryTable(new(Message)).Filter("id", id).RelatedSel(rel...).One(message)
	}

	return message
}

func DeleteMessagesByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Message)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Message)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetMessages(cond *orm.Condition, page, size int, rel ...interface{}) interface{} {
	if size < 1 {
		size = 10
	}

	o := orm.NewOrm()
	messages := make([]*Message, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Message)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(size, size*(page-1)).
			All(&messages)

		data := make(map[string]interface{})
		total, _ := o.QueryTable(new(Message)).SetCond(cond).Filter("hidden", false).Count()
		data["totalCount"] = total
		data["pageSize"] = size
		data["pageIndex"] = page

		if int(total)%size == 0 {
			data["pageCount"] = int(total) / size
		} else {
			data["pageCount"] = int(total)/size + 1
		}

		data["list"] = messages

		return data
	}

	_, _ = o.QueryTable(new(Message)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&messages)
	return messages
}
