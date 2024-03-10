package internal

import (
	"fmt"
	"github.com/RealImage/challenge2016/utils"
)

// AuthorizationMetaData contains the authorization data for a film
type AuthorizationMetaData struct {
	Includes utils.Set
	Excludes utils.Set
	Owner    DistributorID
}

// AuthorizationDB is a map of FilmID to a map of DistributorID to AuthorizationMetaData
var AuthorizationDB map[FilmID]map[DistributorID]AuthorizationMetaData

// NewAuthorizationDB creates a new authorization database
func NewAuthorizationDB() map[FilmID]map[DistributorID]AuthorizationMetaData {
	if AuthorizationDB == nil {
		AuthorizationDB = make(map[FilmID]map[DistributorID]AuthorizationMetaData)
	}
	return AuthorizationDB
}

/*
AuthorizeDistributor authorizes a distributor to distribute a film in a region
- filmID: the ID of the film
- ownerDistributorID: the ID of the distributor who grants permission to agent distributor, "" if the owner is the film producer
- agentDistributorID: the ID of the distributor who is being given permission by the owner distributor
*/
func AuthorizeDistributor(filmID, ownerDistributorID, agentDistributorID string, includes, excludes []string) error {
	if filmID == "" || agentDistributorID == "" {
		return fmt.Errorf("filmID, agentDistributorID cannot be empty")
	}
	if !IsValidFilm(filmID) {
		return fmt.Errorf("invalid filmID")
	}
	if ownerDistributorID != "" && !IsValidDistributor(ownerDistributorID) {
		return fmt.Errorf("invalid ownerDistributorID")
	}
	if !IsValidDistributor(agentDistributorID) {
		return fmt.Errorf("invalid agentDistributorID")
	}
	if len(includes) == 0 {
		return fmt.Errorf("authorization must include at least one region")
	}
	for _, region := range includes {
		if !IsValidRegion(region) {
			return fmt.Errorf("invalid included region: %s", region)
		}
	}
	for _, region := range excludes {
		if !IsValidRegion(region) {
			return fmt.Errorf("invalid excluded region: %s", region)
		}
	}
	// store authorization data if it's valid

	// check if the owner distributor has permission to distribute the film in the regions
	if ownerDistributorID != "" {
		for _, region := range includes {
			if !HasPermission(filmID, ownerDistributorID, region) {
				return fmt.Errorf("owner distributor does not have permission to distribute the film in region: %s", region)
			}
		}
	}
	// Give permission to the agent distributor

	excludedRegionsForAgentDistributor := utils.NewSet()
	excludedRegionsForAgentDistributor.AddItems(excludes)

	excludedRegionsForOwnerDistributor := utils.NewSet()
	if ownerDistributorID != "" {
		excludedRegionsForOwnerDistributor = AuthorizationDB[FilmID(filmID)][DistributorID(ownerDistributorID)].Excludes
		// agent distributor cannot distribute the film in the regions excluded by the owner distributor
		excludedRegionsForAgentDistributor.Union(excludedRegionsForOwnerDistributor)
	}

	includedRegionsForAgentDistributor := utils.NewSet()
	includedRegionsForAgentDistributor.AddItems(includes)

	// Check if the agent distributor already has been given permission to distribute the film
	if _, ok := AuthorizationDB[FilmID(filmID)]; ok {
		if meta, ok := AuthorizationDB[FilmID(filmID)][DistributorID(agentDistributorID)]; ok {
			if len(meta.Includes.Intersection(excludedRegionsForAgentDistributor)) > 0 {
				return fmt.Errorf("agent distributor already has permission to distribute the film in some of the excluded regions")
			}
			if len(meta.Excludes.Intersection(includedRegionsForAgentDistributor)) > 0 {
				return fmt.Errorf("agent distributor doesn't have permission to distribute the film in some of the included regions")
			}
			AuthorizationDB[FilmID(filmID)][DistributorID(agentDistributorID)].Includes.Union(includedRegionsForAgentDistributor)
			AuthorizationDB[FilmID(filmID)][DistributorID(agentDistributorID)].Excludes.Union(excludedRegionsForAgentDistributor)
			return nil
		}
	} else {
		AuthorizationDB[FilmID(filmID)] = make(map[DistributorID]AuthorizationMetaData)
	}
	// film distributor will be given permission for the film for the first time
	AuthorizationDB[FilmID(filmID)][DistributorID(agentDistributorID)] = AuthorizationMetaData{
		Includes: includedRegionsForAgentDistributor,
		Excludes: excludedRegionsForAgentDistributor,
	}
	return nil
}

// HasPermission checks if a distributor has permission to distribute a film in a region
func HasPermission(filmID, distributorID, region string) bool {
	if filmID == "" || distributorID == "" || region == "" {
		fmt.Println("Error: filmID, distributorID, region cannot be empty")
		return false
	}
	if !IsValidFilm(filmID) || !IsValidDistributor(distributorID) || !IsValidRegion(region) {
		fmt.Println("Error: invalid filmID, distributorID, region")
		return false
	}
	// check if the region is excluded for the distributor
	if _, ok := AuthorizationDB[FilmID(filmID)]; ok {
		if _, ok := AuthorizationDB[FilmID(filmID)][DistributorID(distributorID)]; ok {
			if AuthorizationDB[FilmID(filmID)][DistributorID(distributorID)].Excludes.Contains(region) {
				return false
			}
		}
	}
	// check if distributor has permission for the region
	if _, ok := AuthorizationDB[FilmID(filmID)]; ok {
		if _, ok := AuthorizationDB[FilmID(filmID)][DistributorID(distributorID)]; ok {
			if AuthorizationDB[FilmID(filmID)][DistributorID(distributorID)].Includes.Contains(region) {
				return true
			}
		}
	}
	return false
}

// RemoveAuthorizationForFilm deletes the authorization for a film
func RemoveAuthorizationForFilm(filmID string) {
	delete(AuthorizationDB, FilmID(filmID))
}

// RemoveAuthorizationForDistributor deletes the authorization for a distributor to distribute a film
func RemoveAuthorizationForDistributor(filmID, distributorID string) {
	if _, ok := AuthorizationDB[FilmID(filmID)]; ok {
		if _, ok := AuthorizationDB[FilmID(filmID)][DistributorID(distributorID)]; ok {
			delete(AuthorizationDB[FilmID(filmID)], DistributorID(distributorID))
		}
		var allDistributorsIDs []string
		for id := range AuthorizationDB[FilmID(filmID)] {
			allDistributorsIDs = append(allDistributorsIDs, string(id))
		}
		// remove all agent distributors if the owner distributor is being removed
		// NOTE: this will not work if multiple owners are allowed
		for currentDist, meta := range AuthorizationDB[FilmID(filmID)] {
			if meta.Owner == DistributorID(distributorID) {
				delete(AuthorizationDB[FilmID(filmID)], currentDist)
			}
		}
	}
}

// GetAllFilmsAuthorizedForDistributor returns all the films authorized for a distributor
func GetAllFilmsAuthorizedForDistributor(distributorID string) []string {
	var films []string
	for filmID, distributors := range AuthorizationDB {
		for distID := range distributors {
			if string(distID) == distributorID {
				films = append(films, string(filmID))
			}
		}
	}
	return films
}
