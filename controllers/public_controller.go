package controllers

import (
	"antiCov-server/models"
	"antiCov-server/utils"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type PublicController struct {
	BaseController
}

// @description 登录
// @router /login [post]
func (c *PublicController) Login() {

	type postData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	p := new(postData)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, p)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	account := new(models.Account)
	_ = orm.NewOrm().QueryTable(new(models.Account)).Filter("Username", p.Username).One(account)
	if account.Id == 0 {
		c.ReturnSuccess(2, "用户名不存在", nil)
		return
	}

	if utils.Encrypt(p.Password) != account.Password {
		c.ReturnSuccess(2, "密码与该用户名不匹配", nil)
		return
	}

	data := make(map[string]interface{})
	token, err := utils.GenerateTokenString(account.Id)
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
	}
	data["token"] = token
	data["admin_name"] = models.GetAdminByAccountId(account.Id).Name

	c.ReturnSuccess(1, "ok", data)

}

// @description 登录
// @router /admin [get]
func (c *PublicController) GetAdminByToken() {
	id, err := utils.ParseTokenString(c.GetString("token"))
	if err != nil {
		c.ReturnSuccess(2, err.Error(), nil)
		return
	}

	c.ReturnSuccess(1, "ok", models.GetAdminByAccountId(id).Name)
}
