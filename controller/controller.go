package controller

import (
	"net/http"
	datacsv "qube-cinemas-challenge/data-csv"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"Locations": datacsv.Cities})
}