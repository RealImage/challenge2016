package datasource

import (
	"sync"

	"chng2016/pkg/models"
)

// LocalDB ...
type LocalDB interface {
	SetCountryDetails(country *models.Country)
	IsCountryCodeValid(countryCode string) bool
	IsStateCodeValid(stateCode string) bool
	IsCityCodeValid(cityCode string) bool
}

// LocalDBClient ...
type LocalDBClient struct {
	countryCodes map[string]struct{}
	stateCodes   map[string]struct{}
	cityCodes    map[string]struct{}
	sync.RWMutex
}

// NewLocalDBClient ...
func NewLocalDBClient() *LocalDBClient {
	return &LocalDBClient{
		countryCodes: make(map[string]struct{}),
		stateCodes:   make(map[string]struct{}),
		cityCodes:    make(map[string]struct{}),
	}
}
