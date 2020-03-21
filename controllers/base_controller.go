package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 返回结果的数据结构
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许访问方式
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

// Options 允许跨域
func (c *BaseController) TestOptions() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"code": 200, "message": "ok", "data": ""}
	c.ServeJSON()
}

// 返回成功
func (c *BaseController) ReturnSuccess(code int, msg string, data interface{}) {
	if code != 1 {
		fmt.Println(">>>")
		fmt.Println(msg)
	}

	c.AllowCross()
	c.Ctx.Output.SetStatus(200)

	retData := &Result{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.Data["json"] = retData
	c.ServeJSON()
}
