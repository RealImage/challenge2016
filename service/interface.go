package service

import (
	"github.com/challenge2016/models"
	"github.com/gin-gonic/gin"
)

type Service interface{
	AddDistributor(ctx *gin.Context,reqBody *models.Distributor) (*models.Distributor,error)
	GetDistributorByName(ctx *gin.Context,distributorName *string) (*models.Distributor,error)
	CheckDistributorPermission(ctx *gin.Context, reqBody models.CheckPermission) bool
}