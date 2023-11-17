package routes

import (
	"qube-cinemas-challenge/controller"

	"github.com/gin-gonic/gin"
)


func Routes(c *gin.Engine){
	c.GET("/locations", controller.GetLocations)

	//Distributor controller
	c.GET("/distributors", controller.GetDistributor)
	c.POST("/Add-Distributor", controller.AddDistributor)
	c.GET("/get-parent", controller.GetParentDetails)
	c.GET("/get-subDistributors", controller.GetSubDistributors)

	//Permission management controller
	c.GET("/get-included-region", controller.GetIncludedRegion)
	c.POST("/add-included-country", controller.AddIncludedCountry)
	c.POST("/add-included-province", controller.AddIncludedProvince)
	c.POST("/add-included-city", controller.AddIncludedCity)
	c.POST("/remove-included-city", controller.RemoveIncludedCity)
	c.POST("/remove-included-city", controller.RemoveIncludedProvince)
	c.POST("/remove-included-city", controller.RemoveIncludedCountry)

	//Permission checking controller
	c.GET("/city-level-permission", controller.CityLevelPermission)
	c.GET("/province-level-permission", controller.ProvinceLevelPermission)
	c.GET("/country-level-permission", controller.CountryLevelPermission)
}