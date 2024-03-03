package utils

import (
	"fmt"
	"strings"

	"image-challenge/model"
)

// FindParentDistributor finds the parent distributor of a given distributor
// Returns the parent distributor if found, otherwise an empty distributor
func FindParentDistributor(distributor string, distributorMap map[string]map[string]model.Distributor) model.Distributor {
	for _, children := range distributorMap {
		if parent, ok := children[distributor]; ok {
			return parent
		}
	}
	return model.Distributor{}
}

// CheckDistributor checks if a distributor exists and returns its details
// Returns the distributor details and its relation (parent or child) if found, otherwise empty details and relation
func CheckDistributor(distributorMap map[string]map[string]model.Distributor, name string) (model.Distributor, string) {
	if val, ok := distributorMap[name]; ok {
		if distributor, ok := val[name]; ok {
			return distributor, "parent"
		}
	}
	for _, children := range distributorMap {
		if distributor, ok := children[name]; ok {
			return distributor, "child"
		}
	}
	return model.Distributor{}, ""
}

// HasAccess checks if a distributor has access to a particular region
// Returns true if access is granted, otherwise false
func HasAccess(d model.Distributor, region string, type_req string, LocationMap map[string]model.Location) bool {
	location, found := LocationMap[region]
	new_reg1 := ""
	new_reg2 := ""
	new_reg3 := ""

	if !found {
		fmt.Println("entire region not found")
	} else {
		if type_req == "city" {
			new_reg1 = region + "_" + location.State + "_" + location.Country
			new_reg2 = location.State + "_" + location.Country
			new_reg3 = location.Country
		} else if type_req == "state" {
			new_reg1 = location.State + "_" + location.Country
			new_reg2 = location.Country
		} else if type_req == "country" {
			new_reg1 = location.Country
		}
	}
	for _, e := range d.ExcludeRegions {
		if (new_reg1 != "" && strings.HasPrefix(new_reg1, e)) || (new_reg2 != "" && strings.HasPrefix(new_reg2, e)) || (new_reg3 != "" && strings.HasPrefix(new_reg3, e)) {
			return false
		}
	}
	for _, e := range d.IncludeRegions {
		if (new_reg1 != "" && strings.HasPrefix(new_reg1, e)) || (new_reg2 != "" && strings.HasPrefix(new_reg2, e)) || (new_reg3 != "" && strings.HasPrefix(new_reg3, e)) {
			return true
		}
	}

	return false
}
