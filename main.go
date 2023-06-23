package main

import (
	"github.com/challenge2016/http"
	"github.com/challenge2016/models"
	"github.com/challenge2016/preload"
	"github.com/challenge2016/service"
	"github.com/challenge2016/store"
	"github.com/gin-gonic/gin"
)


func main(){

	// initialsing the map
	dMap := models.NewDistributionMaps()

	// preload the data into map 
	// todo - before starting the server , add the filePath
	preload.Preload(dMap,"filePath")

	//log.Println(awsutil.Prettify(dMap.ProvinceMap))

	// dependecy injecttion
	store := store.NewStore(dMap)
	svc := service.New(store)
	http := http.NewHTTP(svc)

	// server initialisation
	router := gin.New()


	// routes 
	router.POST("/addDistributor",http.AddDistributor)
	router.GET("/getDistributor",http.GetDistributorByName)
	router.GET("/checkPermission",http.CheckDistributorPermission)

	// port initialisation
	router.Run(":8080")

}