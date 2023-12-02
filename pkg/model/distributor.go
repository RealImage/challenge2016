package model

import (
	"encoding/json"
)

type AUTH_LEVEL int32

const (
	AUTH_LEVEL_INVALID = iota
	AUTH_LEVEL_COUNTRY
	AUTH_LEVEL_STATE
	AUTH_LEVEL_CITY
)

type Region struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

type Distributor struct {
	Name               string       `json:"name"`
	Authorized         []Region     `json:"authorized_regions,omitempty"`
	UnAuthorized       []Region     `json:"unauthorized_regions,omitempty"`
	AuthorizationLevel AUTH_LEVEL   `json:"authorization_level,omitempty"`
	Parent             *Distributor `json:"-"`
}

type DistributionInfo map[string]DistributorInfo

type DistributorInfo interface {
	GetName() string
	GetAuthorizedRegions() []Region
	GetUnAuthorizedRegions() []Region
	GetParent() *Distributor
	GetAuthLevel() AUTH_LEVEL

	IsParentAuthorizedForRegion(region Region) bool
}

// interface methods implementation
func (d Distributor) GetName() string {
	return d.Name
}

func (d Distributor) GetAuthorizedRegions() []Region {
	return d.Authorized
}

func (d Distributor) GetUnAuthorizedRegions() []Region {
	return d.UnAuthorized
}

func (d Distributor) GetParent() *Distributor {
	return d.Parent
}

func (d Distributor) GetAuthLevel() AUTH_LEVEL {
	return d.AuthorizationLevel
}

func (d Distributor) CheckIfRegionUnAuthorized(region Region) bool {
	for _, unAuthRegion := range d.GetUnAuthorizedRegions() {
		switch {
		case GetAuthLevelForRegion(unAuthRegion) == AUTH_LEVEL_CITY:
			if region == unAuthRegion {
				return false
			}

		case GetAuthLevelForRegion(unAuthRegion) == AUTH_LEVEL_STATE:
			if region.State == unAuthRegion.State && region.Country == unAuthRegion.Country {
				return false
			}

		case GetAuthLevelForRegion(unAuthRegion) == AUTH_LEVEL_COUNTRY:
			if region.Country == unAuthRegion.Country {
				return false
			}
		}
	}

	return true
}

// checks whether distributor's parent is authorized for given region
func (d Distributor) IsParentAuthorizedForRegion(region Region) bool {
	parent := d.GetParent()
	for parent != nil {
		parentAuthLevel := parent.GetAuthLevel()
		regionAuthLevel := GetAuthLevelForRegion(region)
		if parentAuthLevel > regionAuthLevel {
			return false
		}

		if !parent.CheckIfRegionUnAuthorized(region) {
			return false
		}

		parent = parent.GetParent()
	}

	return true
}

// Helper Methods
func GetAuthLevelForRegion(region Region) AUTH_LEVEL {
	switch {
	case region.City != "" && region.State != "" && region.Country != "":
		return AUTH_LEVEL_CITY
	case region.State != "" && region.Country != "":
		return AUTH_LEVEL_STATE
	case region.Country != "":
		return AUTH_LEVEL_COUNTRY
	}

	return AUTH_LEVEL_INVALID
}

func (d Distributor) Marshall() (string, error) {
	serializedInfo, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(serializedInfo), nil
}

// Func validates the given region if the distributor is authorized or not
func (d Distributor) ValidateRegion(region Region) bool {
	distributorAuthLevel := d.GetAuthLevel()
	regionAuthLevel := GetAuthLevelForRegion(region)

	// Validate for authorization level and if region unauthorized
	if distributorAuthLevel > regionAuthLevel || !d.CheckIfRegionUnAuthorized(region) {
		return false
	}

	for _, authRegion := range d.GetAuthorizedRegions() {
		switch {
		case GetAuthLevelForRegion(authRegion) == AUTH_LEVEL_CITY:
			if region == authRegion {
				return true
			}

		case GetAuthLevelForRegion(authRegion) == AUTH_LEVEL_STATE:
			if region.State == authRegion.State && region.Country == authRegion.Country {
				return true
			}

		case GetAuthLevelForRegion(authRegion) == AUTH_LEVEL_COUNTRY:
			if region.Country == authRegion.Country {
				return true
			}
		}
	}

	return false
}
