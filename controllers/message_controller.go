package controllers

import (
	"antiCov-server/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

// @description 创建
// @router /message/create [post]
func (c *PublicController) CreateMessage() {
	//currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))

	message := new(models.Message)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, message)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	// message.Admin = currAdmin

	if _, err = models.CreateMessage(message); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description 获取列表
// @router /messages [get]
func (c *PublicController) GetMessages() {
	//currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))

	page, _ := c.GetInt("page")
	size, _ := c.GetInt("size")

	if size == 0 {
		size = 10
	}

	cond := orm.NewCondition()
	//cond = cond.And("admin_id", currAdmin.Id)

	c.ReturnSuccess(1, "ok", models.GetMessages(cond, page, size))
}

// @description 删除
// @router /message/delete [get]
func (c *PrivateController) DeleteMessage() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))
	_ = currAdmin

	id, err := c.GetInt("id")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if err = models.DeleteMessagesByIds(false, id); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}
