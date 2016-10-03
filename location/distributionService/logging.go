package distributionService

import (
	"time"

	"golang.org/x/net/context"

	"github.com/RealImage/challenge2016/location/domain"
	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) AddDistributor(ctx context.Context, parentDistributorId domain.DistributorId, distributorId domain.DistributorId, locationType domain.LocationType, permission domain.Permission, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "AddLocation",
			"parent_distributor_id", parentDistributorId,
			"distributor_id", distributorId,
			"location_type", locationType,
			"permission", permission,
			"country_code", countryCode,
			"state_code", stateCode,
			"city_code", cityCode,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AddDistributor(ctx, parentDistributorId, distributorId, locationType, permission, countryCode, stateCode, cityCode)
}
