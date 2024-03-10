package internal

import "fmt"

type DistributorID string
type DistributorName string

// distributorDB is a map of DistributorID to distributor
var distributorDB map[DistributorID]DistributorName

// NewDistributorDB creates a new distributor database
func NewDistributorDB() map[DistributorID]DistributorName {
	if distributorDB == nil {
		distributorDB = make(map[DistributorID]DistributorName)
	}
	return distributorDB
}

// AddDistributor adds a new distributor to the database
func AddDistributor(id, name string) error {
	if id == "" || name == "" {
		return fmt.Errorf("invalid distributor data: DistributorID, DistributorName cannot be empty")
	}
	NewDistributorDB()
	if _, ok := distributorDB[DistributorID(id)]; !ok {
		distributorDB[DistributorID(id)] = DistributorName(name)
	} else {
		return fmt.Errorf("distributor already exists")
	}
	return nil
}

// RemoveDistributor removes a distributor from the database
// TODO: remove all authorizations for the distributor
func RemoveDistributor(distributorID string) error {
	if distributorID == "" {
		return fmt.Errorf("invalid distributor data: DistributorID cannot be empty")
	}
	if _, ok := distributorDB[DistributorID(distributorID)]; ok {
		delete(distributorDB, DistributorID(distributorID))
		films := GetAllFilmsAuthorizedForDistributor(distributorID)
		for _, filmID := range films {
			RemoveAuthorizationForDistributor(filmID, distributorID)
		}
	} else {
		return fmt.Errorf("distributor not found")
	}
	return nil
}

// IsValidDistributor checks if the DistributorID is valid
func IsValidDistributor(id string) bool {
	isValid := false
	if id == "" || distributorDB == nil {
		return false
	}
	if _, ok := distributorDB[DistributorID(id)]; ok {
		isValid = true
	}
	return isValid
}
