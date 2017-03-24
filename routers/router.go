package routers

import (
	"challenge2016/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/extract", &controllers.DataExtractController{}, "*:ExtractCities")
}
