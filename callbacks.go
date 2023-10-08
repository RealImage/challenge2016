package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddDistributorPostBody struct {
	ID string `json:"id"`
	Parent *string `json:"parent"`
	Included []string `json:"included"`
	Excluded []string `json:"excluded"`
}

func ListCities(ctx *gin.Context, distributorNetworkObject *DistributorNetwork) {
	SetResponse(ctx, http.StatusOK, distributorNetworkObject.cityData)
}

func ListDistributors(ctx *gin.Context, distributorNetworkObject *DistributorNetwork) {
	SetResponse(ctx, http.StatusOK, distributorNetworkObject.distributorData)
}

func AddDistributors(ctx *gin.Context, distributorNetworkObject *DistributorNetwork) {
	var postData AddDistributorPostBody
	err := ctx.BindJSON(&postData)
	if err != nil {
		SetResponse(ctx, http.StatusBadRequest, err)
		return
	}

	parent := ""
	if postData.Parent != nil {
		parent = *postData.Parent
	}

	err = distributorNetworkObject.AddDistributor(postData.ID, parent, postData.Included, postData.Excluded)
	if err != nil {
		SetResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	SetResponse(ctx, http.StatusOK, "success")
}

func CheckDistributor(ctx *gin.Context, distributorNetworkObject *DistributorNetwork) {
	distributor := ctx.Param("distributor")
	region := ctx.Param("region")
	permitted, err := distributorNetworkObject.IsDistributorPermitted(distributor, region)
	if err != nil {
		SetResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if !permitted {
		SetResponse(ctx, http.StatusForbidden, "NO")
		return
	}
	SetResponse(ctx, http.StatusOK, "YES")
}