package service

import (
	"challenge2016/csvhandler"
	"fmt"
)

type Region struct {
	Country string `json:"country"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
}

type State struct {
	Cities []string
}

type Country struct {
	States map[string]*State
}

type Distributor struct {
	Included map[string]*Country
	Excluded map[string]*Country
	Parent   string
}

type DistributorService struct {
	Distributors map[string]Distributor
	CsvReader    *csvhandler.CsvFileContent
}

func checkIfStateExcluded(state string, country *Country) bool {

	if len(country.States) == 0 {
		return true
	}

	if state, ok := country.States[state]; ok {
		return len(state.Cities) == 0
	}

	return false
}

func checkIfStateIncluded(state string, IncludedStates map[string]*State) bool {

	if len(IncludedStates) == 0 {
		return true
	}

	if _, ok := IncludedStates[state]; ok {

		return true
	}

	return false
}

func checkIfCountryExcluded(country string, excludedCountry map[string]*Country) bool {

	if len(excludedCountry) == 0 {
		return false
	}

	if country, ok := excludedCountry[country]; ok {
		return len(country.States) == 0
	}

	return false
}

func checkIfCountryIncluded(country string, includedCountry map[string]*Country) bool {
	if _, ok := includedCountry[country]; ok {
		return true
	}
	return false
}

// checkCity returns true if inputcity exists in cities list else false
func checkCity(inputCity string, cities []string) bool {
	for _, city := range cities {
		if city == inputCity {
			return true
		}
	}
	return false
}

// AuthorizeSubDistributor authorizes the sub distributor with the given name and list of regions To Be Included and Excluded
func (d *DistributorService) AuthorizeSubDistributor(parentName, name string, regionToIncludes []Region) error {

	if parentName == "" {
		return fmt.Errorf("parent name should be provided")
	}

	dist := Distributor{}

	parentDistributor, ok := d.Distributors[parentName]
	if !ok {

		return fmt.Errorf("No Distributor exists for this parent name: %s", parentName)
	}

	for _, regionToInclude := range regionToIncludes {

		if err := d.CsvReader.ValidateInputRegion(regionToInclude.Country, regionToInclude.State, regionToInclude.City); err != nil {
			return err
		}

		if checkIfCountryExcluded(regionToInclude.Country, parentDistributor.Excluded) {
			return fmt.Errorf("given included country: %s is excluded in parent distributor", regionToInclude.Country)
		}

		if !checkIfCountryIncluded(regionToInclude.Country, parentDistributor.Included) {
			return fmt.Errorf("given included country: %s is not included in parent distributor", regionToInclude.Country)
		}

		// If not included then add the country
		if _, ok := d.Distributors[regionToInclude.Country]; !ok {
			dist.Included = map[string]*Country{
				regionToInclude.Country: &Country{},
			}
		}

		if regionToInclude.State != "" {
			excludedState, excludeCountryExists := parentDistributor.Excluded[regionToInclude.Country]
			includedState, _ := parentDistributor.Included[regionToInclude.Country]

			if excludeCountryExists {
				if checkIfStateExcluded(regionToInclude.State, excludedState) {
					return fmt.Errorf("given included state: %s is excluded in parent distributor", regionToInclude.State)
				}
			}

			if !checkIfStateIncluded(regionToInclude.State, includedState.States) {
				return fmt.Errorf("given included state: %s, is not included in parent distributor", regionToInclude.State)

			}

			stat := dist.Included[regionToInclude.Country]

			stat.States = map[string]*State{
				regionToInclude.State: &State{},
			}

			if regionToInclude.City != "" {

				if excludeCountryExists {

					excludedstate, ok := excludedState.States[regionToInclude.State]
					if ok {
						if len(excludedstate.Cities) == 0 || checkCity(regionToInclude.City, excludedstate.Cities) {
							return fmt.Errorf("given included city: %s, is excluded in parent distributor", regionToInclude.City)
						}

					}
				}

				includedstate, ok := includedState.States[regionToInclude.State]
				if ok {
					if !(len(includedstate.Cities) == 0 || checkCity(regionToInclude.City, includedstate.Cities)) {
						return fmt.Errorf("given included city: %s, is not included in parent distributor", regionToInclude.City)
					}
				}

				stat.States[regionToInclude.State].Cities = append(stat.States[regionToInclude.State].Cities, regionToInclude.City)

			}

			dist.Included[regionToInclude.Country] = stat
		}
	}

	d.Distributors[name] = dist

	return nil
}

// AuthorizeDistributor authorizes the given name and list of regions To Be Included and Excluded
func (d *DistributorService) AuthorizeDistributor(name string, regionToBeIncluded, regionToBeExcluded []Region) error {

	_, ok := d.Distributors[name]
	if ok {
		return fmt.Errorf("Distributor is already authorized %s", name)
	}

	distributor := Distributor{
		Included: make(map[string]*Country),
		Excluded: make(map[string]*Country),
	}

	// Add Included Regions
	for _, region := range regionToBeIncluded {
		if err := d.CsvReader.ValidateInputRegion(region.Country, region.State, region.City); err != nil {
			return err
		}

		// Add the country entry if not existed
		// Else Update existing country object
		_, ok := distributor.Included[region.Country]
		if !ok {
			distributor.Included[region.Country] = &Country{
				States: make(map[string]*State),
			}
		}

		if region.State != "" {
			_, ok := distributor.Included[region.Country].States[region.State]
			if !ok {
				distributor.Included[region.Country].States[region.State] = &State{
					Cities: make([]string, 0),
				}
			}

			if region.City != "" {
				distributor.Included[region.Country].States[region.State].Cities = append(distributor.Included[region.Country].States[region.State].Cities, region.City)
			}
		}
	}

	//Add Excluded Regions
	for _, region := range regionToBeExcluded {
		if err := d.CsvReader.ValidateInputRegion(region.Country, region.State, region.City); err != nil {
			return err
		}

		_, ok := distributor.Excluded[region.Country]
		if !ok {
			distributor.Excluded[region.Country] = &Country{
				States: make(map[string]*State),
			}

		}

		if region.State != "" {
			_, ok := distributor.Excluded[region.Country].States[region.State]
			if !ok {
				distributor.Excluded[region.Country].States[region.State] = &State{
					Cities: make([]string, 0),
				}
			}

			if region.City != "" {
				distributor.Excluded[region.Country].States[region.State].Cities = append(distributor.Excluded[region.Country].States[region.State].Cities, region.City)
			}
		}
	}

	d.Distributors[name] = distributor

	return nil
}

// CheckDistributorPermission checks if a distributor has permission for a specific region.
// It returns true if the distributor is authorized, false if not.
func (d *DistributorService) CheckDistributorPermission(distributorName string, regionToMatch Region) (bool, error) {

	if err := d.CsvReader.ValidateInputRegion(regionToMatch.Country, regionToMatch.State, regionToMatch.City); err != nil {
		return false, err
	}

	// Returns False
	//    1) If all the states under the given country are excluded
	//    2) If all the cities under the given state are excluded
	//    3) If a specific city has been excluded that the user has requested
	if distributor, ok := d.Distributors[distributorName]; ok {
		if country, ok := distributor.Excluded[regionToMatch.Country]; ok {
			if len(country.States) == 0 {
				return false, nil
			} else {
				if state, ok := country.States[regionToMatch.State]; ok {

					if len(state.Cities) == 0 {
						return false, nil
					} else {
						if checkCity(regionToMatch.City, state.Cities) {
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
		if country, ok := distributor.Included[regionToMatch.Country]; ok {
			if len(country.States) == 0 {
				return true, nil
			} else {
				if state, ok := country.States[regionToMatch.State]; ok {

					if len(state.Cities) == 0 {
						return true, nil
					} else {
						if checkCity(regionToMatch.City, state.Cities) {
							return true, nil
						}
					}
				}
			}
		}
	}

	return false, nil
}
