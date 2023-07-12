package handlers

import (
	"chng2016/pkg/datasource"
	"chng2016/pkg/utils"
	"chng2016/pkg/validation"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	HealthHandler(c *gin.Context)
	CheckDistributorPermission(c *gin.Context)
	AddSubDistributor(c *gin.Context)
	AddDistributor(c *gin.Context)
}

type DistributorHandler struct {
	dataStore datasource.Datasource
	validator *validation.Validation
	util      utils.Util
}

func NewDistributorHandler(dataStore datasource.Datasource, validator *validation.Validation, util utils.Util) *DistributorHandler {
	return &DistributorHandler{dataStore: dataStore, validator: validator, util: util}
}
