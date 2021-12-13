package routers

import (
	"challenge2016/controllers"

	"github.com/gin-gonic/gin"
)

func Endpoints(ginrouter *gin.Engine) {
	api := ginrouter.Group("qubecinema")
	api.POST("checkalldistributor", controllers.GetAllDistributors)
	api.POST("insertdistributor", controllers.InsertController)
	api.POST("checkpermission", controllers.CheckDistributorPermissions)
	ginrouter.Run()
}
