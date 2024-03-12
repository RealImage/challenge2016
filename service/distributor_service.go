package service

import (
	"distributor/types"
	"distributor/utils"
	"errors"
)

type DistributorService interface {
	CreateDistributor(req types.DistributorRequest) (types.GenericResponse, error)
	GetDistributorLocationDetails(req types.LocationDetailsReq) (types.GenericResponse, error)
	GetDistributorDetails(req string) (types.GenericResponse, error)
}

type distributorService struct {
}

func (s *distributorService) CreateDistributor(req types.DistributorRequest) (types.GenericResponse, error) {
	var distributor *types.Distributor

	if utils.DistributorsMap[req.DistributorName] == nil {
		distributor = utils.NewDistributor(req.DistributorName, req.ParentDistributorName)
	} else {
		distributor = utils.DistributorsMap[req.DistributorName]
	}
	IncludeLocations := make([]types.LocationIdentifier, 0)
	ExcludeLocations := make([]types.LocationIdentifier, 0)

	for _, include := range req.Includes {
		loc, err := utils.ValidateAndGetLocations(include)
		if err != nil {
			return types.GenericResponse{Message: err.Error()}, err
		}
		IncludeLocations = append(IncludeLocations, loc)
	}

	for _, exclude := range req.Excludes {
		loc, err := utils.ValidateAndGetLocations(exclude)
		if err != nil {
			return types.GenericResponse{Message: err.Error()}, err
		}
		ExcludeLocations = append(ExcludeLocations, loc)
	}

	utils.AddPermission(distributor, IncludeLocations, true)
	utils.AddPermission(distributor, ExcludeLocations, false)

	return types.GenericResponse{Message: "Created Distributor", Data: req}, nil
}

func (s *distributorService) GetDistributorLocationDetails(req types.LocationDetailsReq) (types.GenericResponse, error) {
	if utils.DistributorsMap[req.DistributorName] == nil {
		return types.GenericResponse{Message: "distributor does not exists", Data: req}, errors.New("distributor does not exists")
	}
	distributor := utils.DistributorsMap[req.DistributorName]

	location, err := utils.ValidateAndGetLocations(types.Location{
		CountryCode:  req.CountryCode,
		ProvinceCode: req.ProvinceCode,
		CityCode:     req.CityCode,
	})
	if err != nil {
		return types.GenericResponse{Message: err.Error()}, err
	}

	exists := utils.CheckPermission(distributor, location)
	var msg string
	if exists {
		msg = "Has access"
	} else {
		msg = "Does not have access"
	}
	return types.GenericResponse{Message: msg, Data: req}, nil
}

func (s *distributorService) GetDistributorDetails(name string) (types.GenericResponse, error) {
	if utils.DistributorsMap[name] == nil {
		return types.GenericResponse{Message: "Distributor " + name + " does not exists"}, errors.New("Distributor does not exist")
	}
	return types.GenericResponse{Message: "Distributor details fetched succesfully", Data: utils.DistributorsMap[name]}, nil
}

func NewDistributorService() DistributorService {
	return &distributorService{}
}
