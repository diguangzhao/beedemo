package routers

import (
	"beedemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.TestcController{})
}