package api

import (
	"distribution-mgmnt/app"
	"distribution-mgmnt/internal/svc"
	"distribution-mgmnt/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type DistributionMgnmtServer struct {
}

func NewDistributionMgnmtServer() app.ServerInterface {
	return &DistributionMgnmtServer{}
}

func (dm *DistributionMgnmtServer) GetStatus(c *gin.Context) {
	log.Infoln("started  GetStatus ...")
	c.AbortWithStatusJSON(http.StatusOK, "Distributor Management server is UP  :)")
}

func (dm *DistributionMgnmtServer) AddDistributor(c *gin.Context) {
	log.Infoln("started AddDistributor ..")
	req := app.DistributorDetails{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorln("invalid AddDistributor request ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid AddDistributor details.")
		return
	}
	req = app.DistributorDetails{
		Exclude:           util.ConvertSliceOfStructToUpper(req.Exclude),
		Include:           util.ConvertSliceOfStructToUpper(req.Include),
		Name:              util.RemoveSpacesAndToUpper(req.Name),
		ParentDistributor: util.RemoveSpacesAndToUpper(req.ParentDistributor),
	}
	check := svc.SaveAddDistributor(&req)

	if !check {
		c.AbortWithStatusJSON(http.StatusCreated, "distributor "+req.Name+" is not added")
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, "distributor "+req.Name+" is added")
}
func (dm *DistributionMgnmtServer) GetPermissionsByName(c *gin.Context, distributorName string) {
	log.Infoln("started GetPermissionsByName ..")
	res, ok := svc.GetPermissionsByName(strings.ToUpper(distributorName))
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, "distributor "+distributorName+" not found")
		return
	}
	c.AbortWithStatusJSON(http.StatusFound, res)
}

func (dm *DistributionMgnmtServer) CheckPermission(c *gin.Context) {
	log.Infoln("started CheckPermission ..")
	req := app.CheckPermissionJSONBody{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorln("invalid CheckPermission request ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid CheckPermission details.")
		return
	}
	req = app.CheckPermissionJSONBody{
		Location: app.Location{
			City:     util.RemoveSpacesAndToUpper(req.Location.City),
			Country:  util.RemoveSpacesAndToUpper(req.Location.Country),
			Province: util.RemoveSpacesAndToUpper(req.Location.Province),
		},
		Name: util.RemoveSpacesAndToUpper(req.Name),
	}
	c.AbortWithStatusJSON(http.StatusOK, svc.CheckPermissions(req))
}
