package service

import (
	"challenge2016/csvhandler"
	"challenge2016/model"
	"challenge2016/util"
	"fmt"
)

type Distributor struct {
	Included map[string]*model.Country
	Excluded map[string]*model.Country
	Parent   string
}

type DistributorService struct {
	Distributors map[string]Distributor
	CsvReader    *csvhandler.CsvFileContent
}

// AuthorizeSubDistributor authorizes the sub distributor with the given name and list of regions To Be Included and Excluded
func (d *DistributorService) AuthorizeSubDistributor(parentName, name string, regionToIncludes, regionToExclude []model.Region) error {

	if parentName == "" {
		return fmt.Errorf("parent name should be provided")
	}

	dist := Distributor{}

	parentDistributor, ok := d.Distributors[parentName]
	if !ok {

		return fmt.Errorf("No Distributor exists for this parent name: %s", parentName)
	}

	// Adds the include region if country is included in the parent and not excluded
	for _, regionToInclude := range regionToIncludes {

		country := util.ConvertStringToLowerCase(regionToInclude.Country)
		state := util.ConvertStringToLowerCase(regionToInclude.State)
		city := util.ConvertStringToLowerCase(regionToInclude.City)

		if err := d.CsvReader.ValidateInputRegion(country, state, city); err != nil {
			return err
		}

		if util.CheckIfCountryExcluded(country, parentDistributor.Excluded) {
			return fmt.Errorf("given included country: %s is excluded in parent distributor", country)
		}

		if !util.CheckIfCountryIncluded(country, parentDistributor.Included) {
			return fmt.Errorf("given included country: %s is not included in parent distributor", country)
		}

		// If not included then add the country
		if _, ok := d.Distributors[country]; !ok {
			dist.Included = map[string]*model.Country{
				country: &model.Country{
					States: make(map[string]*model.State),
				},
			}
		}

		// Adds the state, if state is included in parent and not excluded
		if state != "" {
			excludedState, excludeCountryExists := parentDistributor.Excluded[country]
			includedState, _ := parentDistributor.Included[country]

			if excludeCountryExists {
				if util.CheckIfStateExcluded(state, excludedState) {
					return fmt.Errorf("given included state: %s is excluded in parent distributor", state)
				}
			}

			if !util.CheckIfStateIncluded(state, includedState.States) {
				return fmt.Errorf("given included state: %s, is not included in parent distributor", state)

			}

			stat := dist.Included[country]

			if _, ok := stat.States[state]; !ok {
				stat.States[state] = &model.State{}
			}

			// Adds the city if city is included in parent and not excluded
			if city != "" {
				if excludeCountryExists {
					excludedstate, ok := excludedState.States[state]
					if ok {
						if len(excludedstate.Cities) == 0 || util.CheckCityExisted(city, excludedstate.Cities) {
							return fmt.Errorf("given included city: %s, is excluded in parent distributor", city)
						}
					}
				}

				includedstate, ok := includedState.States[state]
				if ok {
					if !(len(includedstate.Cities) == 0 || util.CheckCityExisted(city, includedstate.Cities)) {
						return fmt.Errorf("given included city: %s, is not included in parent distributor", city)
					}
				}

				stat.States[state].Cities = append(stat.States[state].Cities, city)

			}

			dist.Included[country] = stat
		}
	}

	// Adds the exclude region
	for _, regionsToExclude := range regionToExclude {
		country := util.ConvertStringToLowerCase(regionsToExclude.Country)
		state := util.ConvertStringToLowerCase(regionsToExclude.State)
		city := util.ConvertStringToLowerCase(regionsToExclude.City)

		if err := d.CsvReader.ValidateInputRegion(country, state, city); err != nil {
			return err
		}

		// If not included then add the country
		if _, ok := d.Distributors[country]; !ok {
			dist.Excluded = map[string]*model.Country{
				country: &model.Country{
					States: make(map[string]*model.State),
				},
			}
		}

		if state != "" {

			stat := dist.Excluded[country]

			if _, ok := stat.States[state]; !ok {
				stat.States[state] = &model.State{}
			}

			if city != "" {
				stat.States[state].Cities = append(stat.States[state].Cities, city)
			}

			dist.Excluded[country] = stat
		}

	}

	d.Distributors[name] = dist

	return nil
}

// AuthorizeDistributor authorizes the given name and list of regions To Be Included and Excluded
func (d *DistributorService) AuthorizeDistributor(name string, regionToBeIncluded, regionToBeExcluded []model.Region) error {

	_, ok := d.Distributors[name]
	if ok {
		return fmt.Errorf("Distributor is already authorized %s", name)
	}

	distributor := Distributor{
		Included: make(map[string]*model.Country),
		Excluded: make(map[string]*model.Country),
	}

	// Iterate over regions to be included
	for _, region := range regionToBeIncluded {
		country := util.ConvertStringToLowerCase(region.Country)
		state := util.ConvertStringToLowerCase(region.State)
		city := util.ConvertStringToLowerCase(region.City)

		// Validate the incoming region
		if err := d.CsvReader.ValidateInputRegion(country, state, city); err != nil {
			return err
		}

		// Create or update the country object
		if distributor.Included[country] == nil {
			distributor.Included[country] = &model.Country{
				States: make(map[string]*model.State),
			}
		}

		// Create or update the state object
		if region.State != "" {
			stateObj := distributor.Included[country].States[state]
			if stateObj == nil {
				distributor.Included[country].States[state] = &model.State{
					Cities: make([]string, 0),
				}
				stateObj = distributor.Included[country].States[state]
			}

			// Add the city to the state
			if region.City != "" {
				stateObj.Cities = append(stateObj.Cities, city)
			}
		}
	}

	//Add Excluded Regions
	for _, region := range regionToBeExcluded {

		country := util.ConvertStringToLowerCase(region.Country)
		state := util.ConvertStringToLowerCase(region.State)
		city := util.ConvertStringToLowerCase(region.City)

		// validate if the incoming region is a valid one
		if err := d.CsvReader.ValidateInputRegion(country, state, city); err != nil {
			return err
		}

		// Create or update the country object
		if distributor.Excluded[country] == nil {
			distributor.Excluded[country] = &model.Country{
				States: make(map[string]*model.State),
			}
		}

		// Create or update the state object
		if region.State != "" {
			stateObj := distributor.Excluded[country].States[state]
			if stateObj == nil {
				distributor.Excluded[country].States[state] = &model.State{
					Cities: make([]string, 0),
				}
				stateObj = distributor.Excluded[country].States[state]
			}

			// Add the city to the state
			if region.City != "" {
				stateObj.Cities = append(stateObj.Cities, city)
			}
		}
	}

	d.Distributors[name] = distributor

	return nil
}

// CheckDistributorPermission checks if a distributor has permission for a specific region.
// It returns true if the distributor is authorized, false if not.
func (d *DistributorService) CheckDistributorPermission(distributorName string, regionToMatch model.Region) (bool, error) {

	country := util.ConvertStringToLowerCase(regionToMatch.Country)
	state := util.ConvertStringToLowerCase(regionToMatch.State)
	city := util.ConvertStringToLowerCase(regionToMatch.City)

	if err := d.CsvReader.ValidateInputRegion(country, state, city); err != nil {
		return false, err
	}

	// Returns False
	//    1) If all the states under the given country are excluded
	//    2) If all the cities under the given state are excluded
	//    3) If a specific city has been excluded that the user has requested
	if distributor, ok := d.Distributors[distributorName]; ok {
		if country, ok := distributor.Excluded[country]; ok {
			if len(country.States) == 0 {
				return false, nil
			} else {
				if state, ok := country.States[state]; ok {

					if len(state.Cities) == 0 {
						return false, nil
					} else {
						if util.CheckCityExisted(city, state.Cities) {
							return false, nil
						}
					}
				}
			}
		}

		// Returns true
		//    1) If all the states under the given country are included
		//    2) If all the cities under the given state are included
		//    3) If a specific city has been included that the user has requested
		if country, ok := distributor.Included[country]; ok {
			if len(country.States) == 0 {
				return true, nil
			} else {
				if state, ok := country.States[state]; ok {

					if len(state.Cities) == 0 {
						return true, nil
					} else {
						if util.CheckCityExisted(city, state.Cities) {
							return true, nil
						}
					}
				}
			}
		}
	}

	return false, nil
}
