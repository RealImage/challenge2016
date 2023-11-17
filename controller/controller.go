package controller

import (
	"net/http"
	datacsv "qube-cinemas-challenge/data-csv"
	"qube-cinemas-challenge/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"status":true,"Locations": datacsv.Cities})
}

//Distributor Management
func GetDistributor(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"status":true, "distributors":datacsv.Distributor})
}

func AddDistributor(c *gin.Context){
	newDistributor := &models.Distributor{}

	type DistributorData struct{
		Parent string `json:"parent"`
	}
	var newData DistributorData
	if err := c.ShouldBind(&newData);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"error":err.Error()})
		return
	}
	var exist bool
	if newData.Parent != ""{
		for _, distributor:= range datacsv.Distributor{
			exist = distributor.ID==newData.Parent
			if distributor.ID==newData.Parent{
				newDistributor.Parent = distributor
				break
			} 
		}
		if !exist{
			c.JSON(http.StatusUnprocessableEntity, gin.H{"status":false, "message":"Distributor id didn't exist"})
			return
		}
	}
	newDistributor.ID = strconv.Itoa(len(datacsv.Distributor)+1)
	datacsv.Distributor = append(datacsv.Distributor, newDistributor)

	c.JSON(http.StatusAccepted, gin.H{"status":true, "message":"New distributor created with id "+ newDistributor.ID})
}

func GetParentDetails(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	for _,distributor := range datacsv.Distributor{
		if distributor.ID == dist.Id {
			c.JSON(http.StatusAccepted, gin.H{"status":true, "Parent":distributor.Parent})
			return
		}
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{"status":false, "message":"Distributor id doesnot exits"})

}

func GetSubDistributors(c *gin.Context){
	type Dist_Id struct{
		Id string `json:"dist_id"`
	}
	var dist Dist_Id
	if err := c.ShouldBindJSON(&dist);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error":err.Error()})
		return
	}
	var subDistributors []*models.Distributor
	for _,distributor := range datacsv.Distributor {
		if distributor != nil && distributor.Parent != nil && distributor.Parent.ID == dist.Id {
			subDistributors = append(subDistributors, distributor)
		}
	}
	c.JSON(http.StatusAccepted, gin.H{"status":true, "Sub-Distributors":subDistributors})
}