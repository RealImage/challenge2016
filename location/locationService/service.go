package locationService

import (
	"golang.org/x/net/context"

	"github.com/RealImage/challenge2016/location/domain"
)

type Service interface {
	AddLocation(ctx context.Context, countryName string, countryCode string, stateName string, stateCode string, cityName string, cityCode string) (err error)
}

type service struct {
	locationRepository domain.LocationRepository
}

func NewService(locationRepo domain.LocationRepository) *service {
	return &service{
		locationRepository: locationRepo,
	}
}

func (s *service) AddLocation(_ context.Context, countryName string, countryCode string, stateName string, stateCode string, cityName string, cityCode string) (err error) {
	location := &domain.Location{
		CountryName: countryName,
		CountryCode: domain.CountryCode(countryCode),
		StateName:   stateName,
		StateCode:   domain.StateCode(stateCode),
		CityName:    cityName,
		CityCode:    domain.CityCode(cityCode),
	}

	err = location.Validate()
	if err != nil {
		return
	}

	err = s.locationRepository.Store(location)
	return
}
