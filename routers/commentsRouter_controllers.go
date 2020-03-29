package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "GetArticle",
			Router:           `/article`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "CreateArticle",
			Router:           `/article/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "DeleteArticle",
			Router:           `/article/delete`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "UpdateArticle",
			Router:           `/article/update`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "GetArticles",
			Router:           `/articles`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PrivateController"],
		beego.ControllerComments{
			Method:           "DeleteMessage",
			Router:           `/message/delete`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "GetAdminByToken",
			Router:           `/admin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "CreateMessage",
			Router:           `/message/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "GetMessages",
			Router:           `/messages`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "CreateRecord",
			Router:           `/record/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "IncreaseRecord",
			Router:           `/record/increase`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "GetRecords",
			Router:           `/records`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
