package controllers

import (
	"challenge2016/models"
	"challenge2016/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDistributors(c *gin.Context) {
	fmt.Println("inside controller")

	reqDto := models.Distributors{}
	if c.ShouldBind(&reqDto) == nil {
		fmt.Println("binding successful")
	}
	output := service.CheckAllDistributorData()
	c.JSON(http.StatusOK, gin.H{"Values": output})
}

func InsertController(c *gin.Context) {
	fmt.Println("inside controller")
	reqDto := models.Distributors{}
	if c.ShouldBind(&reqDto) == nil {
		fmt.Println("binding successful")
	}
	output, message := service.InsertDistributor(reqDto)
	if output != "" {
		c.JSON(http.StatusOK, gin.H{"Values": output})
	} else {
		c.JSON(http.StatusOK, gin.H{"Values": message})
	}
}

func CheckDistributorPermissions(c *gin.Context) {
	fmt.Println("Inside distributor permissions")
	reqDto := models.Data{}
	if c.ShouldBind(&reqDto) == nil {
		fmt.Println("binding successful")
	}
	output := service.Checkdistributorpermissions(reqDto)
	c.JSON(http.StatusOK, gin.H{"Values": output})
}
