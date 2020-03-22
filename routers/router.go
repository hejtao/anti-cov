package routers

import (
	"antiCov-server/controllers"
	"antiCov-server/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	beego.Router("/*", &controllers.BaseController{}, "options:TestOptions")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/pr",
			beego.NSNamespace("/template",
				beego.NSInclude(
					&controllers.TemplateController{},
				),

				// Todo

			),
		),

		beego.NSNamespace("/pub",
			beego.NSInclude(
				&controllers.PublicController{},
			),
		),
	)

	beego.AddNamespace(ns)

	beego.InsertFilter("/v1/pr/*", beego.BeforeStatic, TokenFilter)
}

// TokenFilter 该过滤器要求请求都带有一个user或者admin的token string
var TokenFilter = func(ctx *context.Context) {
	if ctx.Input.IsOptions() {
		return
	}

	// 解析请求头中的token string
	authString := ctx.Input.Header("Authorization")
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		tokenAuthError(ctx)
		return
	}
	tokenString := kv[1]
	id, err := utils.ParseTokenString(tokenString)
	if err != nil {
		tokenAuthError(ctx)
		return
	}

	ctx.Input.SetData("accountId", id)
}

//Token鉴权失败返回
func tokenAuthError(ctx *context.Context) {
	retData := controllers.Result{
		Code:    3,
		Message: "鉴权失败",
	}
	_ = ctx.Output.JSON(retData, true, false)
}
