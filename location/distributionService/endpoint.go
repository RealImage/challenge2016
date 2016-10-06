package distributionService

import (
	"golang.org/x/net/context"

	"github.com/RealImage/challenge2016/location/domain"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddDistributorEndpoint          endpoint.Endpoint
	CheckLocationPermissionEndpoint endpoint.Endpoint
}

type addDistributorRequest struct {
	ParentDistributorId domain.DistributorId `json:"parent_id"`
	DistributorId       domain.DistributorId `json:"id"`
	LocationType        domain.LocationType  `json:"location_type"`
	Permission          domain.Permission    `json:"permission"`
	CountryCode         domain.CountryCode   `json:"country_code"`
	StateCode           domain.StateCode     `json:"state_code"`
	CityCode            domain.CityCode      `json:"city_code"`
}

type addDistributorResponse struct {
	Err error `json:"error,omitempty"`
}

func (r addDistributorResponse) error() error { return r.Err }

func makeAddDistributorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addDistributorRequest)
		err := s.AddDistributor(ctx, req.ParentDistributorId, req.DistributorId, req.LocationType, req.Permission, req.CountryCode, req.StateCode, req.CityCode)
		return addDistributorResponse{Err: err}, nil
	}
}

func (e Endpoints) AddDistributor(ctx context.Context, parentDistributorId domain.DistributorId, distributorId domain.DistributorId, locationType domain.LocationType, permission domain.Permission, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error) {
	request := addDistributorRequest{
		ParentDistributorId: parentDistributorId,
		DistributorId:       distributorId,
		LocationType:        locationType,
		Permission:          permission,
		CountryCode:         countryCode,
		StateCode:           stateCode,
		CityCode:            cityCode,
	}
	response, err := e.AddDistributorEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(addDistributorResponse).Err
}

type checkLocationPermissionRequest struct {
	DistributorId domain.DistributorId `schema:"id" url:"id"`
	CountryCode   domain.CountryCode   `schema:"country_code" url:"country_code"`
	StateCode     domain.StateCode     `schema:"state_code" url:"state_code"`
	CityCode      domain.CityCode      `schema:"city_code" url:"city_code"`
}

type checkLocationPermissionResponse struct {
	Ok  bool  `json:"ok"`
	Err error `json:"error,omitempty"`
}

func (r checkLocationPermissionResponse) error() error { return r.Err }

func makeCheckLocationPermissionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(checkLocationPermissionRequest)
		ok, err := s.CheckLocationPermission(ctx, req.DistributorId, req.CountryCode, req.StateCode, req.CityCode)
		return checkLocationPermissionResponse{Ok: ok, Err: err}, nil
	}
}

func (e Endpoints) CheckLocationPermission(ctx context.Context, distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (ok bool, err error) {
	request := checkLocationPermissionRequest{
		DistributorId: distributorId,
		CountryCode:   countryCode,
		StateCode:     stateCode,
		CityCode:      cityCode,
	}
	response, err := e.CheckLocationPermissionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(checkLocationPermissionResponse).Ok, response.(checkLocationPermissionResponse).Err
}
