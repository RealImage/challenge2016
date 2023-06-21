package main

import (
	"github.com/challenge2016/models"
	"github.com/challenge2016/preload"
	"github.com/challenge2016/http"
	"github.com/gin-gonic/gin"
)


func main(){

	// initialsing the map
	dMap := models.NewDistributionMaps()

	// todo - take the file path from env file 
	// preload the data into map 
	preload.Preload(dMap,"/Users/abhishekgupta/Desktop/go/src/github.com/challenge2016/cities.csv")

	//log.Println(awsutil.Prettify(dMap.ProvinceMap))

	// dependecy injecttion
	http := http.NewHTTP(dMap)

	// server initialisation
	router := gin.New()


	// routes 
	router.POST("/addDistributor",http.AddDistributor)

	// port initialisation
	router.Run(":8080")

}