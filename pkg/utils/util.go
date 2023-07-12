package utils

import (
	localdb "chng2016/pkg/datasource/localDB"
	"chng2016/pkg/models"
)

// Util ...
type Util interface {
	GetRegionDetails(region string) (*models.Country, error)
	IsValidStateCode(value string) bool
	IsValidCountryCode(value string) bool
	IsValidCityCode(value string) bool
}

// AppUtil ...
type AppUtil struct {
	localDB localdb.LocalDB
}

// NewAppUtil ...
func NewAppUtil(l localdb.LocalDB) *AppUtil {
	return &AppUtil{localDB: l}
}
