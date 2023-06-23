package store

import (
	"github.com/challenge2016/models"
	"github.com/gin-gonic/gin"
)


type Store interface{
	AddDistributor(ctx *gin.Context,reqBody *models.Distributor) *models.Distributor
	GetDistributorByName(ctx *gin.Context, distributorName string) *models.Distributor
	GetLocationDetailsByCity(cityName string) *models.Location
	GetLocationDetailsByProvince(province string) *models.Location
	GetLocationDetailsByCountry(countryName string) *models.Location
}