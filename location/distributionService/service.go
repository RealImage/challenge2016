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

// note: Updating Parent Permission Granted -> Denined will Not Update Permissions to child.
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
				err = s.isCountryHasGrantedPermission(distributorId, d.CountryCode)
				if err != nil {
					return err
				}

				childPerm, _ := s.distribustionRepository.GetCountryPermission(distributorId, d.CountryCode)
				if childPerm != domain.Granted {
					s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.Granted)
					s.copyCountyHirarchi(parentDistributorId, distributorId, d.CountryCode)
				}
			} else {
				s.distribustionRepository.StoreCountry(distributorId, d.CountryCode, domain.Denied)
			}

		case domain.State:
			if d.Permission == domain.Granted {
				err = s.isStateHasGrantedPermission(parentDistributorId, d.CountryCode, d.StateCode)
				if err != nil {
					return err
				}

				childStatePerm, _ := s.distribustionRepository.GetStatePermission(distributorId, d.CountryCode, d.StateCode)
				if childStatePerm != domain.Granted {
					s.storeState(distributorId, d.CountryCode, d.StateCode, domain.Granted)
					s.copyStateHirarchi(parentDistributorId, distributorId, d.CountryCode, d.StateCode)
				}

			} else {
				s.storeState(distributorId, d.CountryCode, d.StateCode, domain.Denied)
			}

		case domain.City:
			if d.Permission == domain.Granted {
				err = s.isCityHasGrantedPermission(parentDistributorId, d.CountryCode, d.StateCode, d.CityCode)
				if err != nil {
					return err
				}

				childPerm, _ := s.distribustionRepository.GetCityPermission(distributorId, d.CountryCode, d.StateCode, d.CityCode)
				if childPerm != domain.Granted {
					s.storeCity(distributorId, d.CountryCode, d.StateCode, d.CityCode, domain.Granted)
				}

			} else {
				s.storeCity(distributorId, d.CountryCode, d.StateCode, d.CityCode, domain.Denied)
			}
		default:
			return domain.ErrInvalidArgument
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
		s.storeState(distributorId, d.CountryCode, d.StateCode, d.Permission)

	case domain.City:
		ok, _ := s.locationRepository.CityExists(d.CountryCode, d.StateCode, d.CityCode)
		if !ok {
			return domain.ErrInvalidArgument
		}

		s.storeCity(distributorId, d.CountryCode, d.StateCode, d.CityCode, d.Permission)
	default:
		return domain.ErrInvalidArgument

	}

	return
}

//storeCity stores city and if intermediate location(county and state) not exists then it adds those location with permission domain.NotDefined
func (s *service) storeCity(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode, cityPermission domain.Permission) {
	_, err := s.distribustionRepository.GetCountryPermission(distributorId, countryCode)
	if err != nil {
		s.distribustionRepository.StoreCountry(distributorId, countryCode, domain.NotDefined)
	}

	_, err = s.distribustionRepository.GetStatePermission(distributorId, countryCode, stateCode)
	if err != nil {
		s.distribustionRepository.StoreState(distributorId, countryCode, stateCode, domain.NotDefined)
	}

	s.distribustionRepository.StoreCity(distributorId, countryCode, stateCode, cityCode, cityPermission)
}

// storeState stores state and if intermediate location(county) not exists then it adds those location with permission domain.NotDefined
func (s *service) storeState(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, statePermission domain.Permission) {
	_, err := s.distribustionRepository.GetCountryPermission(distributorId, countryCode)
	if err != nil {
		s.distribustionRepository.StoreCountry(distributorId, countryCode, domain.NotDefined)
	}

	s.distribustionRepository.StoreState(distributorId, countryCode, stateCode, statePermission)
}

// copyCountyHirarchi copy country hirarchi(states and cities) from parentDistributor to childDistributor
// it also takes care of managing child permission. i.e if child have already denided location permission then it skip that location copy.
func (s *service) copyCountyHirarchi(parentDistributorId domain.DistributorId, distributorId domain.DistributorId, countryCode domain.CountryCode) {
	perentStatePermissions, _ := s.distribustionRepository.ListStatePermission(parentDistributorId, countryCode)
	for _, parentStatePerm := range perentStatePermissions {
		childPerm, _ := s.distribustionRepository.GetStatePermission(distributorId, countryCode, parentStatePerm.StateCode)
		if childPerm == domain.Denied {
			continue
		}
		if childPerm != parentStatePerm.Permission {
			s.distribustionRepository.StoreState(distributorId, countryCode, parentStatePerm.StateCode, parentStatePerm.Permission)
		}

		s.copyStateHirarchi(parentDistributorId, distributorId, countryCode, parentStatePerm.StateCode)
	}
}

// copyStateHirarchi copy country hirarchi(cities) from parentDistributor to childDistributor
// it also takes care of managing child permission. i.e if child have already denided location permission then it skip that location copy.
func (s *service) copyStateHirarchi(parentDistributorId domain.DistributorId, distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode) {
	parentCityPermissions, _ := s.distribustionRepository.ListCityPermission(parentDistributorId, countryCode, stateCode)

	for _, parentCityPerm := range parentCityPermissions {
		childPerm, _ := s.distribustionRepository.GetCityPermission(distributorId, countryCode, stateCode, parentCityPerm.CityCode)

		if childPerm == domain.Denied {
			continue
		}
		if childPerm != parentCityPerm.Permission {
			s.distribustionRepository.StoreCity(distributorId, countryCode, stateCode, parentCityPerm.CityCode, parentCityPerm.Permission)
		}

	}
}

//isCityHasGrantPermission checks if city has granted permission by checking permission of city and if it is not defined then travers to upper hirarchi(state, country) location and check permissions.
func (s *service) isCityHasGrantedPermission(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error) {
	parentCityPerm, err := s.distribustionRepository.GetCityPermission(distributorId, countryCode, stateCode, cityCode)
	if err != nil {
		err = s.isStateHasGrantedPermission(distributorId, countryCode, stateCode)
		if err != nil {
			return err
		}
	} else if parentCityPerm != domain.Granted {
		return domain.ErrParentHaveNotPermission
	}
	return nil
}

func (s *service) isStateHasGrantedPermission(distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode) (err error) {
	parentStatePerm, err := s.distribustionRepository.GetStatePermission(distributorId, countryCode, stateCode)
	if err != nil || parentStatePerm == domain.NotDefined {
		err = s.isCountryHasGrantedPermission(distributorId, countryCode)
		if err != nil {
			return err
		}
	} else if parentStatePerm == domain.Denied {
		return domain.ErrParentHaveNotPermission
	}
	return nil
}

func (s *service) isCountryHasGrantedPermission(distributorId domain.DistributorId, countryCode domain.CountryCode) (err error) {
	parentCountryperm, err := s.distribustionRepository.GetCountryPermission(distributorId, countryCode)
	if err != nil || parentCountryperm != domain.Granted {
		return domain.ErrParentHaveNotPermission
	}
	return nil
}
