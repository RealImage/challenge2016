package service

import (
	"github.com/challenge2016/models"
	"github.com/gin-gonic/gin"
)

type Service interface{
	AddDistributor(ctx *gin.Context,reqBody *models.Distributor)
}