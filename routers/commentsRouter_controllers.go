package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beedemo/controllers:TestcController"] = append(beego.GlobalControllerRouter["beedemo/controllers:TestcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beedemo/controllers:TestcController"] = append(beego.GlobalControllerRouter["beedemo/controllers:TestcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/testc`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beedemo/controllers:TestcController"] = append(beego.GlobalControllerRouter["beedemo/controllers:TestcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/testc/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beedemo/controllers:TestcController"] = append(beego.GlobalControllerRouter["beedemo/controllers:TestcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/testc/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beedemo/controllers:TestcController"] = append(beego.GlobalControllerRouter["beedemo/controllers:TestcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/testc/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
