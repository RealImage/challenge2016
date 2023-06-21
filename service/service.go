package service

import (
	"strings"

	"github.com/challenge2016/models"
	"github.com/gin-gonic/gin"
)

type service struct{
	dMap *models.DistributionMaps
}

func New(dMap *models.DistributionMaps) *service{
	return &service{
		dMap: dMap,
	}
}

func (s *service) AddDistributor(ctx *gin.Context,reqBody *models.Distributor){

	// initilise include fields
	autoInitialiseFields(reqBody.Include,s.dMap)

	// initialise exclude fields
	autoInitialiseFields(reqBody.Exclude,s.dMap)

	// if parentDistributor exist, assign exclude fields also
	if reqBody.ParentDistributor != nil{
		upperCaseName := strings.ToUpper(*reqBody.ParentDistributor)
		reqBody.ParentDistributor = &upperCaseName
	}

	// make store layer call



}

func autoInitialiseFields(loc []models.Location,dMap *models.DistributionMaps){
	for i := range loc{
		if loc[i].City != ""{
			loc[i] = *dMap.CityMap[strings.ToUpper(loc[i].City)]
		}else if loc[i].Country != ""{
			loc[i] = *dMap.CountryMap[strings.ToUpper(loc[i].Country)]
		}else if loc[i].Province != ""{
			loc[i] = *dMap.ProvinceMap[strings.ToUpper(loc[i].Province)]
		}
	}
}