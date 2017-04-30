package routers

import (
	"challenge2016/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DistributorController{}, "*:ListDistributor")
	beego.Router("/new", &controllers.DistributorController{}, "*:NewDistributor")
	beego.Router("/view", &controllers.DistributorController{}, "*:ViewDistributor")
	beego.Router("/new/distributor-provinces", &controllers.DistributorController{}, "*:GetDistributorProvinces")
	beego.Router("/new/distributor-cities", &controllers.DistributorController{}, "*:GetDistributorCities")
}

