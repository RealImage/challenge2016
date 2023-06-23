package http

import (
	"log"
	httpPkg "net/http"

	"github.com/challenge2016/models"
	"github.com/challenge2016/service"
	"github.com/gin-gonic/gin"
)

type http struct{
	svc service.Service
}

func NewHTTP(svc service.Service) *http{
	return &http{
		svc: svc,
	}
}

func (h *http) AddDistributor(ctx *gin.Context) {
	var distributor models.Distributor
	
	err := ctx.ShouldBindJSON(&distributor)
	if err != nil{
		log.Println(err)
		return
	}

	// log.Println(awsutil.Prettify(distributor))

	response,err := h.svc.AddDistributor(ctx,&distributor)
	if err != nil{
		ctx.JSON(httpPkg.StatusBadRequest,gin.H{
			"error":err.Error(),
		})

		return
	}

	ctx.Status(httpPkg.StatusCreated)

	ctx.JSON(httpPkg.StatusCreated,response)
}

func (h *http) GetDistributorByName(ctx *gin.Context) {
	distributorName := ctx.Query("distributor")

	distributor,err := h.svc.GetDistributorByName(ctx,&distributorName)
	if err != nil && err.Error() == "entity not found"{
		ctx.JSON(httpPkg.StatusNotFound,gin.H{
			"error": "distributor not found",
		})

		return
	}

	if err != nil{
		ctx.JSON(httpPkg.StatusNotFound,gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(httpPkg.StatusAccepted,distributor)

	return
}

func (h *http) CheckDistributorPermission(ctx *gin.Context){
	var checkPermission models.CheckPermission
	
	err := ctx.ShouldBindJSON(&checkPermission)
	if err != nil{
		log.Println(err)
		return
	}

	isAllowed := h.svc.CheckDistributorPermission(ctx,checkPermission)
	if !isAllowed{
		ctx.JSON(httpPkg.StatusBadRequest,gin.H{
			"message": "no",
		})

		return
	}

	ctx.JSON(httpPkg.StatusAccepted,gin.H{
		"message": "yes",
	})	
}