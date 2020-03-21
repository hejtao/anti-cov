package controllers

import (
	"antiCov-server/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type TemplateController struct {
	BaseController
}

// @description 创建
// @router /template/create [post]
func (c *TemplateController) CreateTemplate() {
	currAdmin := c.Ctx.Input.GetData("currAdmin").(*models.Admin)
	_ = currAdmin

	template := new(models.Template)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, template)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	cond := orm.NewCondition()
	if models.TemplateExist(cond) {
		c.ReturnSuccess(2, "请勿重复录入", nil)
		return
	}

	if _, err = models.CreateTemplate(template); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description 获取列表
// @router /templates [get]
func (c *TemplateController) GetTemplates() {
	currAdmin := c.Ctx.Input.GetData("currAdmin").(*models.Admin)
	_ = currAdmin

	page, err := c.GetInt("page")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	cond := orm.NewCondition()
	c.ReturnSuccess(1, "ok", models.GetTemplates(cond, page))
}

// @description 删除
// @router /templates/delete [get]
func (c *TemplateController) DeleteTemplate() {
	currAdmin := c.Ctx.Input.GetData("currAdmin").(*models.Admin)
	_ = currAdmin

	id, err := c.GetInt("id")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if err = models.DeleteTemplatesByIds(false, id); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description
// @router /template [get]
func (c *TemplateController) GetTemplate() {
	currAdmin := c.Ctx.Input.GetData("currAdmin").(*models.Admin)
	_ = currAdmin

	id, err := c.GetInt("id")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", models.GetTemplateById(id))
}

// @description 更新
// @router /template/update [put]
func (c *TemplateController) UpdateTemplate() {
	currAdmin := c.Ctx.Input.GetData("currAdmin").(*models.Admin)
	_ = currAdmin

	template := new(models.Template)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, template)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if err = models.UpdateTemplate(template, ""); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}
