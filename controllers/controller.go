package controllers

import (
	"challenge2016/models"
	"net/http"
	"github.com/gin-gonic/gin"
)
 
func GetAll(c *gin.Context) {
        response := models.DistributorResponse{
            Distributors: models.DistributerList,
        }
    
        c.JSON(http.StatusOK, response)
}
        
   

func CheckPermissions(c *gin.Context) {
    distributorName := c.Query("distributor_name")
    location := c.Query("location")
    //check if asked location is a valid location
    valid:=models.CheckCode(location)
    if valid{
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
    }else{
    c.JSON(http.StatusOK, gin.H{"error": "Not a valid location"})
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
