package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	distributorNetworkObject := NewDistributorNetwork()
	
	router := gin.Default()

	router.GET("/list-cities", func(ctx *gin.Context) {
		ListCities(ctx, distributorNetworkObject)
	})

	router.POST("/add-distributors", func (ctx *gin.Context) {
		AddDistributors(ctx, distributorNetworkObject)
	})

	router.GET("/list-distributors", func (ctx *gin.Context) {
		ListDistributors(ctx, distributorNetworkObject)
	})

	router.GET("/check-distributor-region/:distributor/:region", func(ctx *gin.Context) {
		CheckDistributor(ctx, distributorNetworkObject)
	})

	router.Run(":8000")
}