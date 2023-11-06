package controllers

import (
	"challenge2016/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{"error": "Hello Up"})
}

func CheckPermissions(c *gin.Context) {
    distributorName := c.Query("distributor_name")
    location := c.Query("location")
	var distributer_detail models.Distributor 
	for _,d:=range models.DistributerList{
		if d.Name==distributorName{
		distributer_detail=d
		break
		}
}
    status := models.IsPermitted(distributer_detail, location)

    if status {
        c.JSON(http.StatusOK, gin.H{"message": "YES"})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "NO"})
    }
}

func AddDistributor(c *gin.Context) {
    var distributorInput models.DistributorInput

    if err := c.ShouldBindJSON(&distributorInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	newDistributor := models.Distributor{
		Name:       distributorInput.Name,
		Cities:     distributorInput.Cities,
		States:     distributorInput.States,
		Countries:  distributorInput.Countries,
		ExCities:   distributorInput.ExCities,
		ExStates:   distributorInput.ExStates,
		ExCountries: distributorInput.ExCountries,
		Addedby: 	distributorInput.Addedby,
	}
    create_status := models.CreateDistributor(newDistributor)

    if create_status == false {
        c.JSON(http.StatusForbidden, gin.H{"message": "Not permitted"})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "Distributor added successfully"})
    }
}
