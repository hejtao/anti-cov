package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"],
        beego.ControllerComments{
            Method: "GetTemplate",
            Router: `/template`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"],
        beego.ControllerComments{
            Method: "CreateTemplate",
            Router: `/template/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"],
        beego.ControllerComments{
            Method: "UpdateTemplate",
            Router: `/template/update`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"],
        beego.ControllerComments{
            Method: "GetTemplates",
            Router: `/templates`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"] = append(beego.GlobalControllerRouter["antiCov-server/controllers:TemplateController"],
        beego.ControllerComments{
            Method: "DeleteTemplate",
            Router: `/templates/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
