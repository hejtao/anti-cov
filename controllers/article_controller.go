package controllers

import (
	"antiCov-server/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

// @description 创建
// @router /article/create [post]
func (c *PrivateController) CreateArticle() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))

	article := new(models.Article)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, article)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	article.AdminId = currAdmin.Id

	if _, err = models.CreateArticle(article); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description 获取列表
// @router /articles [post]
func (c *PrivateController) GetArticles() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))
	_ = currAdmin

	type postDate struct {
		Page     int   `json:"page"`
		Sections []int `json:"sections"`
	}

	p := new(postDate)
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, p)

	cond := orm.NewCondition()
	if len(p.Sections) > 0 {
		cond = cond.And("Section__in", p.Sections)
	}

	c.ReturnSuccess(1, "ok", models.GetArticles(cond, p.Page))
}

// @description 删除
// @router /article/delete [get]
func (c *PrivateController) DeleteArticle() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))
	_ = currAdmin

	id, err := c.GetInt("id")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if err = models.DeleteArticlesByIds(false, id); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}

// @description
// @router /article [get]
func (c *PrivateController) GetArticle() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))
	_ = currAdmin

	id, err := c.GetInt("id")
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", models.GetArticleById(id))
}

// @description 更新
// @router /article/update [put]
func (c *PrivateController) UpdateArticle() {
	currAdmin := models.GetAdminByAccountId(c.Ctx.Input.GetData("accountId").(int))
	_ = currAdmin

	article := new(models.Article)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, article)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	if err = models.UpdateArticle(article); err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", nil)
}
