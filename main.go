package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var Locations LocationData
	Locations.Init()
	Locations.Distributors = make(map[string]DistributorData)
	Locations.DistributorParent = make(map[string]string)

	router := gin.Default()

	router.POST("/create-distributor", Locations.CreateDistributor)

	router.GET("/distributors", Locations.GetDistributors)

	router.GET("/permission-check/:distributor/:permission", Locations.DistributorPermissionCheck)

	router.Run(":8090")
}
