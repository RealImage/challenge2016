package handlers

import (
	"challenge2016/model"
	service "challenge2016/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorizeDistributor struct {
	Name            string         `json:"name"`
	IncludedRegions []model.Region `json:"includedregions"`
	ExcludedRegions []model.Region `json:"excludedregions"`
	ParentName      string         `json:"parentname,omitempty"`
}
type Handler struct {
	Distributor service.DistributorService
}

func (h *Handler) AuthorizeDistributor(c *gin.Context) {

	var authDistributor authorizeDistributor

	if err := c.BindJSON(&authDistributor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Distributor.AuthorizeDistributor(authDistributor.Name, authDistributor.IncludedRegions, authDistributor.ExcludedRegions); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, authDistributor)
}

func (h *Handler) CheckDistributorAccess(c *gin.Context) {
	country := c.Query("country")
	state := c.Query("state")
	city := c.Query("city")

	name := c.Param("distributor")

	regionToMatch := model.Region{
		Country: country,
		State:   state,
		City:    city,
	}

	exists, err := h.Distributor.CheckDistributorPermission(name, regionToMatch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		c.IndentedJSON(http.StatusOK, "No")
		return
	}

	c.IndentedJSON(http.StatusOK, "Yes")
}

func (h *Handler) AuthorizeSubDistributor(c *gin.Context) {

	var authDistributor authorizeDistributor

	if err := c.BindJSON(&authDistributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Distributor.AuthorizeSubDistributor(authDistributor.ParentName, authDistributor.Name, authDistributor.IncludedRegions, authDistributor.ExcludedRegions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, authDistributor)

}
