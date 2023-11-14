package authorisation

import (
	"challengeQube/dtos"
	"challengeQube/internal/globals"
	"context"
	"errors"
	"fmt"
	"log"
)

type Authorize struct {
	allow bool
}

func New(allow bool) *Authorize {
	return &Authorize{
		allow: allow,
	}
}

func (a *Authorize) GiveAuthorisation(ctx context.Context, req *dtos.AuthorisationReq) error {
	//basic checks

	if req == nil {
		return errors.New("request can't be nil")
	}

	if req.DistributorName == "" {
		return errors.New("distributor name can't be empty")
	}
	if len(req.Include) == 0 {
		return errors.New("included list can't be empty")
	}

	if _, avail := globals.DistributorData[req.DistributorName]; avail {
		return errors.New("distributor name already exists")
	}

	if req.ParentName != "" {
		if _, avail := globals.DistributorData[req.ParentName]; !avail {
			return errors.New("this parent name doesn't exist in master data")
		}
	}

	distributor := dtos.Distributor{
		Included: make(map[string]*dtos.Country),
		Excluded: make(map[string]*dtos.Country),
	}

	allowedInAllRegionInc := make(map[string]bool, 0)
	excludeInAllRegion := make(map[string]bool, 0)

	for ind, inc := range req.Include {
		// handling of subset entry. we can return error also insted of continue
		if allowedInAllRegionInc[inc.Country] {
			continue
		}
		if allowedInAllRegionInc[inc.State] {
			continue
		}
		// handling of duplicated entry
		err := checkForDuplicateData(ctx, distributor, inc, globals.TypeInclude)
		if err != nil {
			log.Println("error: ", err)
			return errors.New("included array has duplicated data")
		}

		// this function validate the country, state and city codes with the master data
		err = validateRegion(ctx, inc, globals.TypeInclude, ind)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
		if inc.State == "" && inc.City == "" {
			allowedInAllRegionInc[inc.Country] = true
		}
		if inc.State != "" && inc.City == "" {
			allowedInAllRegionInc[inc.State] = true
		}
		if req.ParentName != "" {
			// this function validate the request such that the region/location data in it is subset of the parent
			_, err := validateChildWithParentRegion(ctx, inc, req.ParentName, globals.TypeInclude, ind)
			if err != nil {
				log.Println("error: ", err)
				return err
			}
		}
		// this adds the data to the distributor struct
		addAuthorisationData(ctx, inc, distributor.Included)
	}

	for ind, exclude := range req.Exclude {
		if excludeInAllRegion[exclude.Country] {
			continue
		}
		if excludeInAllRegion[exclude.State] {
			continue
		}
		err := checkForDuplicateData(ctx, distributor, exclude, globals.TypeExclude)
		if err != nil {
			log.Println("error: ", err)
			return errors.New("excluded array has duplicated data")
		}
		// this checks for duplicate data among exclude and include both
		err = checkForDuplicateData(ctx, distributor, exclude, globals.TypeInclude)
		if err != nil {
			log.Println("error: ", err)
			return errors.New("excluded and included array both have same data")
		}

		err = validateRegion(ctx, exclude, globals.TypeExclude, ind)
		if err != nil {
			log.Println("error: ", err)
			return err
		}

		if exclude.State == "" && exclude.City == "" {
			allowedInAllRegionInc[exclude.Country] = true
		}
		if exclude.State != "" && exclude.City == "" {
			allowedInAllRegionInc[exclude.State] = true
		}

		addAuthorisationData(ctx, exclude, distributor.Excluded)
	}

	distributor.Parent = req.ParentName
	if globals.DistributorData == nil {
		globals.DistributorData = make(map[string]dtos.Distributor, 0)
	}
	globals.DistributorData[req.DistributorName] = distributor
	return nil
}

func checkForDuplicateData(ctx context.Context, distributor dtos.Distributor, loc *dtos.Location, reqType string) error {
	var data map[string]*dtos.Country

	switch reqType {
	case globals.TypeInclude:
		data = distributor.Included
	case globals.TypeExclude:
		data = distributor.Excluded
	}

	if _, exists := data[loc.Country]; exists {
		if loc.Country == "" && len(distributor.Excluded[loc.Country].States) == 0 {
			return errors.New("exclude has repeated data")
		}
		if _, exists := data[loc.Country].States[loc.State]; exists {
			if loc.City == "" && len(data[loc.Country].States[loc.State].Cities) == 0 {
				return errors.New("exclude has repeated data")
			}
			if _, exists := data[loc.Country].States[loc.State].Cities[loc.City]; exists {
				return errors.New("the data is duplicated in excluded")
			}
		}
	}
	return nil
}

func addAuthorisationData(ctx context.Context, data *dtos.Location, res map[string]*dtos.Country) {

	if res[data.Country] == nil {
		res[data.Country] = &dtos.Country{
			States: make(map[string]*dtos.State),
		}
	}

	if data.State != "" {
		stateObj := res[data.Country].States[data.State]
		if stateObj == nil {
			res[data.Country].States[data.State] = &dtos.State{
				Cities: make(map[string]bool, 0),
			}
			stateObj = res[data.Country].States[data.State]
		}

		// Add the city to the state
		if data.City != "" {
			cityMap := make(map[string]bool, 0)
			cityMap[data.City] = true
			stateObj.Cities = cityMap
		}
	}
}

func validateChildWithParentRegion(ctx context.Context, region *dtos.Location, parentName string, reqFrom string, ind int) (bool, error) {

	if _, exists := globals.DistributorData[parentName]; !exists {
		return false, errors.New("parent doesn't exists")
	}
	data, dataBool := globals.DistributorData[parentName].Included[region.Country]
	if reqFrom == globals.TypeExclude {
		data, dataBool = globals.DistributorData[parentName].Excluded[region.Country]
	}
	if exists := dataBool; !exists && region.Country != "" {
		return false, errors.New("country is not subset of parent")
	}

	if region.State != "" {
		//case of first time
		var newState bool
		if len(data.States) == 0 {
			err := validateRegion(ctx, region, reqFrom, ind)
			if err != nil {
				log.Println("error: ", err)
				return false, err
			}
			newState = true
		}
		if _, exists := data.States[region.State]; !newState && !exists {
			return false, errors.New("country is not subset of parent")
		}
	}

	if region.City != "" {
		//case of first time
		var newCity bool
		if len(data.States[region.State].Cities) == 0 {
			err := validateRegion(ctx, region, reqFrom, ind)
			if err != nil {
				log.Println("error: ", err)
				return false, err
			}
			newCity = true
		}
		if _, exists := data.States[region.State].Cities[region.City]; !newCity && !exists {
			return false, errors.New("country is not subset of parent")
		}
	}

	return true, nil
}

func validateRegion(ctx context.Context, region *dtos.Location, reqFrom string, ind int) error {
	if region.Country != "" {
		if _, exists := globals.MasterData[region.Country]; !exists {
			return errors.New("region's country in " + reqFrom + " at position " + fmt.Sprint(ind+1) + " doesn't exist in master data")
		}
	} else {
		return errors.New("region's country is empty in " + reqFrom + " at position " + fmt.Sprint(ind+1))
	}

	if region.State != "" {
		if _, exists := globals.MasterData[region.Country].States[region.State]; !exists {
			return errors.New("region's state in " + reqFrom + " at position " + fmt.Sprint(ind+1) + " doesn't exist in master data")
		}
	}

	if region.City != "" {
		if region.State != "" {
			if _, exists := globals.MasterData[region.Country].States[region.State].Cities[region.City]; !exists {
				return errors.New("region's city in " + reqFrom + " at position " + fmt.Sprint(ind+1) + " doesn't exist in master data")
			}
		} else {
			return errors.New("region's state is empty in " + reqFrom + " at position " + fmt.Sprint(ind+1))
		}
	}

	return nil
}
