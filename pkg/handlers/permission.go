package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistributorPermissionRequest struct {
	CountryCode   string `json:"countryCode" validate:"required,isValidCountryCode"`
	StateCode     string `json:"stateCode" validate:"required,isValidStateCode"`
	CityCode      string `json:"cityCode" validate:"required,isValidCityCode"`
	DistributorID string `json:"distributorID" validate:"required"`
}

// CheckDistributorPermission ...
func (d *DistributorHandler) CheckDistributorPermission(c *gin.Context) {
	req := DistributorPermissionRequest{
		CountryCode:   c.Query("countryCode"),
		StateCode:     c.Query("stateCode"),
		CityCode:      c.Query("cityCode"),
		DistributorID: c.Param("distributorID"),
	}
	// validating input arguments
	if err := d.validator.Validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": d.validator.CustomValidationError(&req, err),
		})
		return
	}

	distributor := d.dataStore.GetCache(req.DistributorID)
	if distributor == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "distributor not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hasPermission": d.dataStore.HasPermission(distributor.TrieRoot, distributor.Permissions.Include, distributor.Permissions.Exclude, req.CountryCode, req.StateCode, req.CityCode),
	})
}

// func (d *DistributorHandler) isDistributorPermitted(distributorID, countryCode, stateCode, cityCode string) bool {
// 	permission := d.distributorCache.GetCache(distributorID)
// 	isPermitted := false
// 	if permission != nil {
// 		p := permission.(*models.PermissionDetails)
// 		if distributorDetails, isDistributorIDFound := p.Permission[distributorID]; isDistributorIDFound {
// 			if countryDetails, isCountryFound := distributorDetails[countryCode]; isCountryFound {
// 				if stateDetails, isStateFound := countryDetails.StateWisePermission[stateCode]; isStateFound {
// 					_, isCityExcluded := stateDetails.Excluded[cityCode]
// 					if !isCityExcluded && !stateDetails.IsCompletelyExcluded {
// 						isPermitted = true
// 					}
// 				} else {
// 					if countryDetails.IsCompletelyIncluded {
// 						isPermitted = true
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return isPermitted
// }
