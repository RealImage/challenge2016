package utility

import (
	"fmt"
	"strings"

	"realImage.com/m/model"
)

// Helper function to find parent Distributor of a given distributor
func FindParentDistributor(distributor string, distributorMap map[string][]model.Distributor) model.Distributor {

	// Iterate through all the children in the map, if not found, return empty struct
	for _, val := range distributorMap {
		for _, childD := range val {
			if childD.Name == distributor {
				return val[0]
			}
		}
	}
	return model.Distributor{}
}

// Helper function to return Distributor{} object for a given distributor name
func CheckDistributor(distributorMap map[string][]model.Distributor, name string) (model.Distributor, string) {

	var val []model.Distributor
	var ok bool
	if val, ok = distributorMap[name]; ok && val[0].Name == name {
		return val[0], "parent"
	}

	for _, ele := range distributorMap {
		for _, childDistributor := range ele {
			if childDistributor.Name == name {
				return childDistributor, "child"
			}
		}
	}
	return model.Distributor{}, ""
}

// Helper function to check if a distributor has access to a particular location/region
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
