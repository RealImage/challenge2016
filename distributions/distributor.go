package distributions /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	disterr "../err"
)

//Distributor Struct stores the Name, Permission Template, Parent Distributor Deatils of a Distributor
type Distributor struct {
	Name        string
	permissions CountryMap
	parent      *Distributor
}

//Location Struct stores the CityCode, ProvinceCode, CountryCode Received as the Input
type Location struct {
	City     string `json:"cityCode"`
	Province string `json:"provinceCode"`
	Country  string `json:"countryCode"`
}

// Initialize intializes a Distributor
func (distributor *Distributor) Initialize(name string, parent *Distributor) {
	distributor.Name = name
	distributor.parent = parent
	distributor.permissions = CreateCSVTemplate()
}

//Include the Locations to Distributor Template
func (distributor *Distributor) Include(locations []Location) error {
	//if the distributor is not a sub-distributor
	if distributor.parent == nil {
		return distributor.permissions.Include(locations)
	}
	//else if the distributor is a sub-distributor(has parent)
	if distributor.parent.permissions.IsAllowed(locations) {
		return distributor.permissions.Copy(locations, distributor.parent.permissions)
	}
	return disterr.DistributionError()
}

//Exclude the Locations from Distributor Template
func (distributor *Distributor) Exclude(locations []Location) error {
	return distributor.permissions.Exclude(locations)
}

//VerifyPermissions for the given locations for the Specified Distributor
func (distributor *Distributor) VerifyPermissions(location Location) bool {
	return distributor.permissions.Verify(location)
}
