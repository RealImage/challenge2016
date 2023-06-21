package http

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/challenge2016/models"
	"github.com/challenge2016/service"
	"github.com/gin-gonic/gin"

)

type http struct{
	dMap *models.DistributionMaps
	svc service.Service
}

func NewHTTP(dMap *models.DistributionMaps) *http{
	return &http{
		dMap: dMap,
	}
}

func (h *http) AddDistributor(ctx *gin.Context) {
	var distributor models.Distributor
	
	err := ctx.ShouldBindJSON(&distributor)
	if err != nil{
		log.Println(err)
		return
	}

	log.Println(awsutil.Prettify(distributor))

	h.svc.AddDistributor(ctx,&distributor)
}