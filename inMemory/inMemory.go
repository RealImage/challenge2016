package inmemory

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/nikhilsiwach28/Cinema-Distribution-System/models"
)

type InMemory interface {
	Add(models.Distributor) bool
	Remove(uuid.UUID) bool
	CheckIfCountryAuth(uuid.UUID, string) bool
	CheckIfStateAuth(uuid.UUID, string, string) bool
	CheckIfCityAuth(uuid.UUID, string, string, string) bool
	CheckIfIncludesAuthByParent(*models.Distributor, uuid.UUID) error
	CheckIfExcludesAuthByParent(*models.Distributor, uuid.UUID) error
}

type inMemory struct {
	distributors []models.Distributor
}

func (q *inMemory) Add(distributor models.Distributor) bool {

	q.distributors = append(q.distributors, distributor)
	return true
}

func (q *inMemory) Remove(ID uuid.UUID) bool {

	for i, dis := range q.distributors {
		if dis.ID == ID {
			q.distributors = append(q.distributors[:i], q.distributors[i+1:]...)
			return true
		}
	}
	return false
}

func (q *inMemory) CheckIfCountryAuth(id uuid.UUID, country string) bool {
	for _, distributor := range q.distributors {
		if distributor.ID == id {
			for _, includedCountry := range distributor.Access.Include {
				if includedCountry == country {
					return true
				}
			}
		}
	}
	return false
}

func (q *inMemory) CheckIfStateAuth(id uuid.UUID, state, country string) bool {
	for _, distributor := range q.distributors {
		if distributor.ID == id {
			for _, included := range distributor.Access.Include {
				if included == state || included == country {
					for _, excluded := range distributor.Access.Exclude {
						if excluded == state {
							return false
						}
					}
					return true
				}
			}
		}
	}
	return false
}

func (q *inMemory) CheckIfCityAuth(id uuid.UUID, city, state, country string) bool {
	for _, distributor := range q.distributors {
		if distributor.ID == id {
			for _, included := range distributor.Access.Include {
				if included == city || included == state || included == country {
					for _, excluded := range distributor.Access.Exclude {
						if excluded == state || excluded == city {
							return false
						}
					}
					return true
				}
			}
		}
	}
	return false
}

func (q *inMemory) findDistributorByID(ID uuid.UUID) *models.Distributor {
	for _, distributor := range q.distributors {
		if distributor.ID == ID {
			// Found a matching ID, return the distributor
			return &distributor
		}
	}
	// ID not found
	return nil
}

func (q *inMemory) CheckIfIncludesAuthByParent(dist *models.Distributor, parentId uuid.UUID) error {
	parent := q.findDistributorByID(parentId)
	if ok := checkIfParentIncludesOrExcludes(dist.Access.Include, parent.Access.Include); ok != true {
		return errors.New("Parent Not Authorised for the Regions")
	}
	return nil
}

func (q *inMemory) CheckIfExcludesAuthByParent(dist *models.Distributor, parentId uuid.UUID) error {
	parent := q.findDistributorByID(parentId)
	if ok := checkIfParentIncludesOrExcludes(dist.Access.Exclude, parent.Access.Exclude); ok != true {
		return errors.New("Parent Not Authorised for the Regions")
	}
	return nil
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

// Utility

func checkIfParentIncludesOrExcludes(distRegions, parentRegions []string) bool {

	// Loop through distributor's includes or excludes and check if they are all present in parent's includes or excludes
	for _, location := range distRegions {
		city, state, country := splitLocation(location)

		// Check if this location (city/state/country) is not present in the parent's includes as well as excludes
		if !includesLocation(city, state, country, parentRegions) {
			return false
		}
	}

	return true
}

func splitLocation(location string) (city, state, country string) {
	// Implement your logic to split the location string
	// For example:
	parts := strings.Split(location, "-")
	switch len(parts) {
	case 1:
		country = parts[0]
	case 2:
		state = parts[0]
		country = parts[1]
	case 3:
		city = parts[0]
		state = parts[1]
		country = parts[2]
	}
	return city, state, country
}

func includesLocation(city, state, country string, includes []string) bool {
	for _, loc := range includes {
		locCity, locState, locCountry := splitLocation(loc)

		// Check if the location (city/state/country) matches any of the includes
		if locCity == city && locState == state && locCountry == country {
			return true
		}
	}
	return false
}
