package routers

import (
	"challenge2016/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DistributorController{}, "*:ListDistributor")
	beego.Router("/new", &controllers.DistributorController{}, "*:NewDistributor")
}

