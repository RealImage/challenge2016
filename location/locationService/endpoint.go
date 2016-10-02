package locationService

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddLocationEndpoint endpoint.Endpoint
}

type addLocationRequest struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	StateName   string `json:"state_name"`
	StateCode   string `json:"state_code"`
	CityName    string `json:"city_name"`
	CityCode    string `json:"city_code"`
}

type addLocationResponse struct {
	Err error `json:"error,omitempty"`
}

func (r addLocationResponse) error() error { return r.Err }

func makeAddLocationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addLocationRequest)
		err := s.AddLocation(ctx, req.CountryName, req.CountryCode, req.StateName, req.StateCode, req.CityName, req.CityCode)
		return addLocationResponse{Err: err}, nil
	}
}

func (e Endpoints) AddLocation(ctx context.Context, countryName string, countryCode string, stateName string, stateCode string, cityName string, cityCode string) (err error) {
	request := addLocationRequest{
		CountryName: countryName,
		CountryCode: countryCode,
		StateName:   stateName,
		StateCode:   stateCode,
		CityName:    cityName,
		CityCode:    cityCode,
	}
	response, err := e.AddLocationEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return response.(addLocationResponse).Err
}
