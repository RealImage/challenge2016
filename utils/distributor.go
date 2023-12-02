package utils

import (
	"challenge2016/pkg/model"
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

// Func Parses the given input csv records and returns a distribution Info map
func ParseCsv(
	log *slog.Logger,
	inputData [][]string,
) (model.DistributionInfo, error) {

	distributionInfo := make(model.DistributionInfo)
	for index, record := range inputData {
		if index == 0 || len(record) == 0 {
			/* Ignore the CSV headers and empty records */
			continue
		}

		name, include, exclude := record[0], record[1], record[2]
		if include == exclude {
			log.Error("invalid testcase, include and exclude regions are same, skipping it.",
				"row:", record,
			)
			continue
		}

		// validate if given name is a sub-distributor format
		if !validateDistributorChain(log, name, distributionInfo) {
			continue
		}

		includeRegion, excludeRegion := getIncludeExcludeRegions(include, exclude)
		distributorInfo, err := getDistributorInfo(
			log,
			name,
			includeRegion,
			excludeRegion,
			distributionInfo,
		)
		if err != nil {
			log.Error("invalid distributor data, skipping it",
				"name", name,
				"error:", err,
			)
			removeFromDistribution(log, distributorInfo, distributionInfo)
			continue
		}

		log.Debug("Updating distribution info for given distributor",
			"name", distributorInfo.GetName(),
		)
		// update the distribution info map
		updateDistributionInfo(log, distributorInfo, distributionInfo)
	}

	return distributionInfo, nil
}

// Func does the validation of the given distributor csv input and returns the distributor info
func getDistributorInfo(
	log *slog.Logger,
	distributorName string,
	includeRegion []string,
	excludeRegion []string,
	distributionInfo model.DistributionInfo,
) (model.DistributorInfo, error) {
	var distributor model.Distributor

	// check if given distributor is a sub-distributor
	isSubDistributor := strings.Contains(distributorName, "<")
	if isSubDistributor {
		distributorChain := strings.Split(distributorName, "<")
		subDistributorName, parentDistributorName := distributorChain[0], distributorChain[1]

		// Update the parent pointer for each distributor correctly
		if parentDistributor, ok := distributionInfo[parentDistributorName].(model.Distributor); ok {
			log.Debug("Updating parent pointer for distributor",
				"Sub Distributor: ", subDistributorName,
				"parent Distributor: ", parentDistributor,
			)
			distributor.Parent = &parentDistributor
		}

		distributor.Name = subDistributorName
	} else {
		distributor.Name = distributorName
	}

	/* get the include regions details*/
	if err := validateAndAuthorizeRegion(
		log,
		&distributor,
		includeRegion,
		isSubDistributor,
		distributionInfo,
	); err != nil {
		return distributor, err
	}

	/* get the exclude regions details */
	if err := validateAndUnAuthorizeRegion(
		log,
		&distributor,
		excludeRegion,
		isSubDistributor,
		distributionInfo,
	); err != nil {
		return distributor, err
	}
	return distributor, nil
}

// Func gets the include and exclude regions for given csv input
func getIncludeExcludeRegions(include string, exclude string) ([]string, []string) {
	var includeRegion, excludeRegion []string
	if include != "" {
		includeRegion = strings.Split(include, "-")
	}
	if exclude != "" {
		excludeRegion = strings.Split(exclude, "-")
	}

	return includeRegion, excludeRegion
}

// Func validates the given distributor's include region
// and adds the regions to authorized regions list
func validateAndAuthorizeRegion(
	log *slog.Logger,
	distributor *model.Distributor,
	includeRegion []string,
	isSubDistributor bool,
	distributionInfo model.DistributionInfo,
) error {
	if len(includeRegion) == 0 {
		return nil
	}
	region, regionLevel, err := getRegionDetails(log, includeRegion)
	if err != nil {
		return err
	}

	// authorizationLevel := getAuthorizationLevel(distributor, regionLevel)
	authorizationLevel := getAuthLevelBasedOnRegion(log, distributor, distributionInfo, regionLevel)
	if authorizationLevel == model.AUTH_LEVEL_INVALID {
		return errors.New("could not authorize region, authorization error")
	}
	distributor.AuthorizationLevel = authorizationLevel

	if isSubDistributor && !distributor.IsParentAuthorizedForRegion(region) {
		return errors.New("failed parent authorization check")
	} else {
		distributor.Authorized = append(distributor.Authorized, region)
		log.Debug("Adding region to authorized list")
	}

	return nil
}

// Func gets the authorization level based on given region auth level.
// Should only be called when we are trying to authorize a region for the given distributor
func getAuthLevelBasedOnRegion(
	log *slog.Logger,
	d *model.Distributor,
	disInfo model.DistributionInfo,
	regionAuthLevel model.AUTH_LEVEL,
) model.AUTH_LEVEL {
	var authLevel = regionAuthLevel

	if _, ok := disInfo[d.Name]; ok {
		oldAuthLevel := disInfo[d.Name].GetAuthLevel()
		authLevel = max(authLevel, oldAuthLevel)
	}

	isSubDistributor := (d.Parent != nil)
	if isSubDistributor && authLevel <= d.Parent.AuthorizationLevel {
		return model.AUTH_LEVEL_INVALID
	}

	return authLevel
}

// Func validates the given distributor's exclude region
// and adds the regions to unauthorized regions list
func validateAndUnAuthorizeRegion(
	log *slog.Logger,
	distributor *model.Distributor,
	excludeRegion []string,
	isSubDistributor bool,
	distributionInfo model.DistributionInfo,
) error {
	if len(excludeRegion) == 0 {
		return nil
	}

	region, regionLevel, err := getRegionDetails(log, excludeRegion)
	if err != nil {
		return err
	}

	// get the authorization level from the current input (if valid)
	// or from the distribution info map (if present)
	var authorizationLevel = distributor.AuthorizationLevel
	if authorizationLevel == model.AUTH_LEVEL_INVALID {
		if oldInfo, ok := distributionInfo[distributor.Name]; ok {
			authorizationLevel = oldInfo.GetAuthLevel()
		}
	}

	if authorizationLevel == model.AUTH_LEVEL_INVALID {
		log.Error("distributor is not authorized for any region",
			"name", distributor.Name,
		)
		return fmt.Errorf("failed unauthorized check")
	}

	// To UnAuthorize a region for given distributor, it makes sense only to un-authorize
	// region which is below the authorization level
	// Ex: if authorization is set a country level, he can be unauthorized for state or city regions only.
	if regionLevel <= distributor.AuthorizationLevel {
		// log as error and ignore this region
		log.Error("could not unauthorize region, redundant unauthorization",
			"distributor", distributor.Name,
			"Region", region,
			"Region level:", regionLevel,
			"Authorization level: ", distributor.AuthorizationLevel,
		)
		return nil
	}
	distributor.UnAuthorized = append(distributor.UnAuthorized, region)
	log.Debug("Adding region to unauthorized list")

	return nil
}

// Func transforms the input csv region and returns it along with the region Authorization Level
func getRegionDetails(log *slog.Logger, region []string) (model.Region, model.AUTH_LEVEL, error) {
	var city, state, country string
	var regionAuthLevel model.AUTH_LEVEL = model.AUTH_LEVEL_INVALID

	switch len(region) {
	case 3:
		city, state, country = region[0], region[1], region[2]
		regionAuthLevel = model.AUTH_LEVEL_CITY
	case 2:
		state, country = region[0], region[1]
		regionAuthLevel = model.AUTH_LEVEL_STATE
	case 1:
		country = region[0]
		regionAuthLevel = model.AUTH_LEVEL_COUNTRY
	default:
		errString := fmt.Sprintf("INVALID input region, "+
			"should be <city>-<state>-<country> format region: %v\n", region)

		return model.Region{}, regionAuthLevel, fmt.Errorf(errString)
	}

	var transformedRegion = model.Region{
		City:    city,
		State:   state,
		Country: country,
	}
	return transformedRegion, regionAuthLevel, nil
}

// Func validates the distributor chain, if the given input name is a sub-distributor
// It checks the existence of parent distributors in the distribution Info map
func validateDistributorChain(
	log *slog.Logger,
	name string,
	distributionInfo model.DistributionInfo,
) bool {
	isSubDistributor := strings.Contains(name, "<")

	if isSubDistributor {
		/* Validate the distributor chain input */
		distributorChain := strings.Split(name, "<")
		for i := len(distributorChain) - 1; i > 0; i-- {
			parentDistributor := distributorChain[i]

			/* If parent distributor itself is not present, mark as invalid*/
			if _, ok := distributionInfo[parentDistributor]; !ok && parentDistributor != "" {
				log.Error("parent distributor not found in map, skipping it",
					"parent name", parentDistributor,
					"distributor chain", name,
				)
				return false
			}
		}
	}

	return true
}

// Func updates the distribution Info map for given distributor Info
func updateDistributionInfo(
	log *slog.Logger,
	distributorInfo model.DistributorInfo,
	distributionInfo model.DistributionInfo,
) {
	name := distributorInfo.GetName()

	if _, ok := distributionInfo[name]; !ok {
		distributionInfo[name] = distributorInfo
	} else {
		dInfo := distributionInfo[name]

		authorized, unauthorized := dInfo.GetAuthorizedRegions(), dInfo.GetUnAuthorizedRegions()

		authorized = append(authorized, distributorInfo.GetAuthorizedRegions()...)
		unauthorized = append(unauthorized, distributorInfo.GetUnAuthorizedRegions()...)
		authLevel := max(dInfo.GetAuthLevel(), distributorInfo.GetAuthLevel())
		parent := distributorInfo.GetParent()

		distributionInfo[name] = model.Distributor{
			Name:               name,
			Authorized:         authorized,
			UnAuthorized:       unauthorized,
			AuthorizationLevel: authLevel,
			Parent:             parent,
		}
	}
}

// Func removes the given distributor from the distribution Info map
// and also resets the parent pointers for the distributors appropriately
func removeFromDistribution(
	log *slog.Logger,
	distributorInfo model.DistributorInfo,
	distributionInfo model.DistributionInfo,
) {
	distributor := distributorInfo.(model.Distributor)
	for _, d := range distributionInfo {
		parent := distributionInfo[d.GetName()].GetParent()
		if distributor.GetParent() == parent {
			parent = nil
		}
	}

	delete(distributionInfo, distributorInfo.GetName())
}

// Validation Methods (will be called from main)

// Func check if the given distributors names exist in the distribution map
func checkIfDistributorExists(
	log *slog.Logger,
	distributionInfo model.DistributionInfo,
	distributor ...string,
) bool {
	for _, name := range distributor {
		if _, ok := distributionInfo[name]; !ok {
			log.Error("Distributor does not exist in the distribution map", "name", name)
			return false
		}
	}

	return true
}

// Func validates the given input and returns whether given distributor is authorized
// for given region or not
func ValidateInput(
	log *slog.Logger,
	name string,
	inputRegion string,
	distributionInfo model.DistributionInfo,
) bool {
	var distributorInfo model.DistributorInfo
	var distributorName string

	isSubDistributor := strings.Contains(name, "<")
	if isSubDistributor {
		chain := strings.Split(name, "<")
		subDistributor, parentDistributor := chain[0], chain[1]

		if !checkIfDistributorExists(
			log,
			distributionInfo,
			subDistributor,
			parentDistributor,
		) {
			return false
		}

		distributorName = subDistributor
	} else {
		if !checkIfDistributorExists(log,
			distributionInfo,
			name,
		) {
			return false
		}

		distributorName = name
	}

	distributorInfo = distributionInfo[distributorName]
	distributor := distributorInfo.(model.Distributor)

	validateRegion, _ := getIncludeExcludeRegions(inputRegion, "")
	region, _, err := getRegionDetails(log, validateRegion)
	if err != nil {
		return false
	}

	return distributor.ValidateRegion(region)
}
