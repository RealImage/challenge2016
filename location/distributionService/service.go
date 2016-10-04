package distributionService

import (
	"golang.org/x/net/context"

	"github.com/RealImage/challenge2016/location/domain"
)

type Service interface {
	AddDistributor(ctx context.Context, parentDistributorId domain.DistributorId, distributorId domain.DistributorId, locationType domain.LocationType, permission domain.Permission, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error)
}

type service struct {
	distribustionRepository domain.DistributionRepository
	locationRepository      domain.LocationRepository
}

func NewService(distribustionRepo domain.DistributionRepository, locationRepo domain.LocationRepository) *service {
	return &service{
		distribustionRepository: distribustionRepo,
		locationRepository:      locationRepo,
	}
}

func (s *service) AddDistributor(_ context.Context, parentDistributorId domain.DistributorId, distributorId domain.DistributorId, locationType domain.LocationType, permission domain.Permission, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error) {
	d := &domain.DistributorPermission{
		LocationType: locationType,
		Permission:   permission,
		CountryCode:  countryCode,
		StateCode:    stateCode,
		CityCode:     cityCode,
	}

	err = d.Validate()
	if err != nil {
		return
	}

	if len(parentDistributorId) > 1 {
		switch d.LocationType {
		case domain.Country:
			if d.Permission == domain.Granted {
				parentCountryperm, err := s.distribustionRepository.GetCountryPermission(parentDistributorId, d.CountryCode)
				if err != nil {
					return domain.ErrParentHaveNotPermission
				}

				if parentCountryperm == domain.Granted {
					//TODO Copy whole county map insted of assining granted perm
					s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.Granted)
				} else {
					return domain.ErrParentHaveNotPermission
				}
			} else {
				s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.Denied)
			}

		}
		return
	}

	switch d.LocationType {
	case domain.Country:
		ok, _ := s.locationRepository.CountryExists(d.CountryCode)
		if !ok {
			return domain.ErrInvalidArgument
		}

		s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, d.Permission)
	case domain.State:
		ok, _ := s.locationRepository.StateExists(d.CountryCode, d.StateCode)
		if !ok {
			return domain.ErrInvalidArgument
		}

		perm, err := s.distribustionRepository.GetCountryPermission(distributorId, d.CountryCode)
		if err != nil {
			s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.NotDefined)
		}
		if perm == domain.Denied {
			break
		}
		s.distribustionRepository.StoreState(distributorId, d.CountryCode, d.StateCode, d.Permission)

	case domain.City:
		ok, _ := s.locationRepository.CityExists(d.CountryCode, d.StateCode, d.CityCode)
		if !ok {
			return domain.ErrInvalidArgument
		}

		perm, err := s.distribustionRepository.GetCountryPermission(distributorId, d.CountryCode)
		if err != nil {
			s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.NotDefined)
		}

		if perm == domain.Denied {
			break
		}

		perm, err = s.distribustionRepository.GetStatePermission(distributorId, d.CountryCode, d.StateCode)
		if err != nil {
			s.distribustionRepository.StoreState(distributorId, d.CountryCode, d.StateCode, domain.NotDefined)
		}

		if perm == domain.Denied {
			break
		}

		s.distribustionRepository.StoreCity(distributorId, d.CountryCode, d.StateCode, d.CityCode, d.Permission)
	default:
		return domain.ErrInvalidArgument

	}

	return
}
