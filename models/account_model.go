package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Account struct {
	Id int `json:"id"`

	Username string `json:"username"`
	Password string `json:"password"`

	Hidden     bool      `json:"hidden,omitempty"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"update_time" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Account))
}

func CreateAccount(account *Account) (int, error) {
	id, err := orm.NewOrm().Insert(account)
	return int(id), err
}

func UpdateAccount(account *Account, cols ...string) error {
	_, err := orm.NewOrm().Update(account, cols...)
	return err
}

// 如果返回的 account.Id == 0, 则出错
func GetAccountById(id int, rel ...interface{}) *Account {
	o := orm.NewOrm()
	account := new(Account)

	if len(rel) == 0 {
		_ = o.QueryTable(new(Account)).Filter("id", id).One(account)
	} else {
		_ = o.QueryTable(new(Account)).Filter("id", id).RelatedSel(rel...).One(account)
	}

	return account
}

func DeleteAccountsByIds(hard bool, ids ...int) error {
	if hard {
		_, err := orm.NewOrm().QueryTable(new(Account)).Filter("Id__in", ids).Delete()
		return err
	}

	_, err := orm.NewOrm().QueryTable(new(Account)).Filter("Id__in", ids).Update(orm.Params{
		"Hidden": true,
	})
	return err
}

func GetAccounts(cond *orm.Condition, page int, rel ...interface{}) interface{} {
	o := orm.NewOrm()
	accounts := make([]*Account, 0)

	if page > 0 {
		_, _ = o.QueryTable(new(Account)).
			SetCond(cond).Filter("hidden", false).
			RelatedSel(rel...).
			OrderBy("-Id").
			Limit(10, 10*(page-1)).
			All(&accounts)

		data := make(map[string]interface{})
		data["count"], _ = o.QueryTable(new(Account)).SetCond(cond).Filter("hidden", false).Count()
		data["accounts"] = accounts

		return data
	}

	_, _ = o.QueryTable(new(Account)).
		SetCond(cond).Filter("hidden", false).
		RelatedSel(rel...).
		OrderBy("-Id").
		All(&accounts)
	return accounts
}

func AccountExist(cond *orm.Condition) bool {
	return orm.NewOrm().QueryTable(new(Account)).SetCond(cond).Exist()
}
