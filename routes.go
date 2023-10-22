package main

import (
	"github.com/gin-gonic/gin"
)


// Initialize the Gin router and API route handlers
func SetupRouter(regions []Region) *gin.Engine {
    r := gin.Default()

    r.POST("/distributor", CreateDistributor(regions))
	r.GET("/distributors", GetAllDistributors)
	r.GET("/distributor/:distributor/region/:region", CheckDistributorPermission)

    return r
}