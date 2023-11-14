package authorisation

import (
	"challengeQube/dtos"
	"challengeQube/internal/globals"
	"context"
	"errors"
	"log"
)

func (a *Authorize) GetAuthorisationStatus(ctx context.Context, req *dtos.GetStatusReq) (bool, error) {
	if req.DistributorName == "" {
		return false, errors.New("distributor name can't be empty")
	}
	region := req.Region
	if region.Country == "" {
		return false, errors.New("country name can't be empty")
	}

	err := validateRegion(ctx, &region, "", 0)
	if err != nil {
		log.Println("error: ", err)
		return false, err
	}
	if val, exists := globals.DistributorData[req.DistributorName]; exists {
		countryExists, stateExists := false, false
		//check exclude case
		if _, exists := val.Excluded[region.Country]; exists {
			countryExists = true
		}
		if countryExists {
			if _, exists := val.Excluded[region.Country].States[region.State]; region.State != "" && exists {
				stateExists = true
			}
		}
		// edge case handling: ind-up in exclude and ind include and req is for ind or vice versa maybe
		if stateExists && countryExists && region.City == "" {
			return false, nil
		}

		if stateExists {
			if _, exists := val.Excluded[region.Country].States[region.State].Cities[region.City]; region.City != "" && exists {
				return false, nil
			}
		}
		//check include
		if _, exists := val.Included[region.Country]; exists {
			if len(val.Included[region.Country].States) == 0 {
				return true, nil
			}
		} else {
			return false, nil
		}

		if _, exists := val.Included[region.Country].States[region.State]; region.State != "" && exists {
			if len(val.Included[region.Country].States[region.State].Cities) == 0 {
				return true, nil
			}
		} else {
			return false, nil
		}
		if _, exists := val.Included[region.Country].States[region.State].Cities[region.City]; region.City != "" && exists {
			return true, nil
		} else {
			return false, nil
		}

	} else {
		return false, errors.New("no such distributor exists in master data")
	}
}
