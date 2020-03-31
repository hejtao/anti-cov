package controllers

import (
	"antiCov-server/models"
	"github.com/astaxie/beego/orm"
)

// @description 获取列表
// @router /records [get]
func (c *PublicController) GetRecords() {
	records := make([]models.Record, 0)

	_, _ = orm.NewOrm().QueryTable(new(models.Record)).
		All(&records, "Count", "Type", "Lang", "Section", "SecLang")

	c.ReturnSuccess(1, "ok", records)
}

// @description
// @router /record/increase [get]
func (c *PublicController) IncreaseRecord() {
	cond := orm.NewCondition()

	t, _ := c.GetInt("type")
	l, _ := c.GetInt("lang")
	s, _ := c.GetInt("sec")
	sl, _ := c.GetInt("sec_lang")

	cond = cond.And("Type", t)
	cond = cond.And("Lang", l)
	cond = cond.And("Section", s)
	cond = cond.And("SecLang", sl)

	if !orm.NewOrm().QueryTable(new(models.Record)).SetCond(cond).Exist() {
		record := &models.Record{
			Count:   1,
			Type:    t,
			Lang:    l,
			Section: s,
			SecLang: sl,
		}
		recordId, _ := models.CreateRecord(record)

		_, _ = models.CreateTimeline(&models.Timeline{
			RecordId: recordId,
		})

		c.ReturnSuccess(1, "ok", nil)
		return
	}

	_, _ = orm.NewOrm().QueryTable(new(models.Record)).SetCond(cond).Update(orm.Params{
		"count": orm.ColValue(orm.ColAdd, 1),
	})

	record := new(models.Record)
	_ = orm.NewOrm().QueryTable(new(models.Record)).SetCond(cond).One(record, "Id")

	_, _ = models.CreateTimeline(&models.Timeline{
		RecordId: record.Id,
	})

	c.ReturnSuccess(1, "ok", nil)
}
