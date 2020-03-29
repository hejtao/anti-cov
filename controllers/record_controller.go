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

	if orm.NewOrm().QueryTable(new(models.Record)).Filter("Type", record.Type).Exist() {
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

	page, err := c.GetInt("page")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", models.GetRecords(orm.NewCondition(), page))
}

// @description
// @router /record/increase [get]
func (c *PublicController) IncreaseRecord() {

	t, err := c.GetInt("type")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if _, err := orm.NewOrm().QueryTable(new(models.Record)).Filter("Type", t).Update(orm.Params{
		"count": orm.ColValue(orm.ColAdd, 1),
	}); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}
