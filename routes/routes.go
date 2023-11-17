package routes

import (
	"qube-cinemas-challenge/controller"

	"github.com/gin-gonic/gin"
)


func Routes(c *gin.Engine){
	c.GET("/locations", controller.GetLocations)
	c.GET("/distributors", controller.GetDistributor)
	c.POST("/Add-Distributor", controller.AddDistributor)
	c.GET("/get-parent", controller.GetParentDetails)
	c.GET("/get-subDistributors", controller.GetSubDistributors)

	c.GET("/get-included-region", controller.GetIncludedRegion)
	c.POST("/add-included-country", controller.AddIncludedCountry)
}