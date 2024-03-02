package service

import (
	"errors"
	"strings"

	"github.com/saurabh-sde/challenge2016_saurabh/model"
	"github.com/saurabh-sde/challenge2016_saurabh/utils"
)

func GetDistributors() map[string]model.Distributor {
	// utils.Println(utils.DistributorMap)
	return utils.DistributorMap
}

func AddDistributor(req *utils.NewDistributorRequest) (*model.Distributor, error) {
	// check code present into sheet
	isValidLoc := ValidateLocation(req.Includes, req.Excludes)
	if !isValidLoc {
		utils.Error("Invalid Location")
		err := errors.New("invalid Location")
		return nil, err
	}
	// add distributor
	distributor, ok := utils.DistributorMap[req.Name]
	if ok {
		// update  distibutor
		utils.Println("Updating Distributor")
		if len(distributor.Includes) == 0 {
			distributor.Includes = make([]string, 0)
		}
		distributor.Includes = append(distributor.Includes, req.Includes...)
		if len(distributor.Excludes) == 0 {
			distributor.Excludes = make([]string, 0)
		}
		distributor.Excludes = append(distributor.Excludes, req.Excludes...)
	} else {
		// add distributor
		utils.Println("Adding new Distributor")
		if len(req.Includes) == 0 {
			req.Includes = make([]string, 0)
		}
		if len(req.Excludes) == 0 {
			req.Excludes = make([]string, 0)
		}
		distributor = model.Distributor{
			Name:              req.Name,
			Includes:          req.Includes,
			Excludes:          req.Excludes,
			ParentDistributor: req.Parent,
		}
	}
	distributor.Includes = utils.RemoveDuplicateStr(distributor.Includes)
	distributor.Excludes = utils.RemoveDuplicateStr(distributor.Excludes)
	utils.DistributorMap[req.Name] = distributor
	utils.Println("DistributorMap:", utils.DistributorMap)
	return &distributor, nil
}

func GetDistributorPermission(name string, locations []string) (reslt map[string]string, err error) {
	reslt = make(map[string]string, len(locations))
	distributor, ok := utils.DistributorMap[name]
	if !ok {
		utils.Error("Error wrong distributor name")
		err := errors.New("invalid Distributor")
		return reslt, err
	}

	for _, l := range locations {
		reslt[l] = CheckDistributorPermissionForLocation(distributor.Name, l)
	}
	return reslt, nil
}

func CheckDistributorPermissionForLocation(name, loc string) string {
	distributor, ok := utils.DistributorMap[name]
	if !ok {
		return "NO"
	}
	utils.Println("CheckDistributorPermissionForLocation: ", distributor)
	for _, i := range distributor.Excludes {
		if strings.HasSuffix(loc, i) {
			return "NO"
		}
	}

	for _, i := range distributor.Includes {
		if strings.Contains(loc, i) && distributor.ParentDistributor == "" {
			return "YES"
		}
	}

	if distributor.ParentDistributor != "" {
		return CheckDistributorPermissionForLocation(distributor.ParentDistributor, loc)
	}
	return "NO"
}
