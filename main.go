package main

import (
	"challenge2016/csvhandler"
	"challenge2016/handlers"
	"challenge2016/service"

	"github.com/gin-gonic/gin"
)

func main() {

	csvFileContent := csvhandler.CsvFileContent{
		CityMap: map[string]*csvhandler.Country{},
	}

	if err := csvFileContent.LoadCsv(); err != nil {
		panic(err)
	}

	r := gin.Default()

	distributorMap := service.DistributorService{
		Distributors: make(map[string]service.Distributor),
		CsvReader:    &csvFileContent,
	}

	h := &handlers.Handler{
		Distributor: distributorMap,
	}

	// Register api routes
	r.POST("/authdistributor", h.AuthorizeDistributor)

	r.GET("/distributors/:distributor/access", h.CheckDistributorAccess)

	r.POST("/authsubdistributor", h.AuthorizeSubDistributor)

	if err := r.Run(); err != nil {
		panic(err)
	}

}
