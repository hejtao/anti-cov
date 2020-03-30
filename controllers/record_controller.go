package controllers

import (
	"antiCov-server/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

// @description 创建
// @router /record/create [post]
func (c *PublicController) CreateRecord() {

	record := new(models.Record)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, record)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	cond := orm.NewCondition()
	if record.Type > 0 {
		cond = cond.And("Type", record.Type)
	}

	if record.Lang > 0 {
		cond = cond.And("Lang", record.Lang)
	}

	if record.Section > 0 {
		cond = cond.And("Section", record.Section)
	}

	if record.SecLang > 0 {
		cond = cond.And("SecLang", record.SecLang)
	}

	if orm.NewOrm().QueryTable(new(models.Record)).SetCond(cond).Exist() {
		c.ReturnSuccess(2, "请勿重复创建", nil)
		return
	}

	if _, err = models.CreateRecord(record); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description 获取列表
// @router /records [get]
func (c *PublicController) GetRecords() {
	cond := orm.NewCondition()
	if t, _ := c.GetInt("type"); t > 0 {
		cond = cond.And("Type", t)
	}

	if l, _ := c.GetInt("lang"); l > 0 {
		cond = cond.And("Lang", l)
	}

	if s, _ := c.GetInt("sec"); s > 0 {
		cond = cond.And("Section", s)
	}

	if sl, _ := c.GetInt("sec_lang"); sl > 0 {
		cond = cond.And("SecLang", sl)
	}

	records := make([]models.Record, 0)
	_, _ = orm.NewOrm().QueryTable(new(models.Record)).
		SetCond(cond).All(&records, "Count", "Type", "Lang", "Section", "SecLang")

	c.ReturnSuccess(1, "ok", records)
}

// @description
// @router /record/increase [get]
func (c *PublicController) IncreaseRecord() {

	cond := orm.NewCondition()
	if t, _ := c.GetInt("type"); t > 0 {
		cond = cond.And("Type", t)
	}

	if l, _ := c.GetInt("lang"); l > 0 {
		cond = cond.And("Lang", l)
	}

	if s, _ := c.GetInt("sec"); s > 0 {
		cond = cond.And("Section", s)
	}

	if sl, _ := c.GetInt("sec_lang"); sl > 0 {
		cond = cond.And("SecLang", sl)
	}

	if _, err := orm.NewOrm().QueryTable(new(models.Record)).SetCond(cond).Update(orm.Params{
		"count": orm.ColValue(orm.ColAdd, 1),
	}); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}
