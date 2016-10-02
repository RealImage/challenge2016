package locationService

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) AddLocation(ctx context.Context, countryName string, countryCode string, stateName string, stateCode string, cityName string, cityCode string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "AddLocation",
			"country_name", countryName,
			"country_code", countryCode,
			"state_name", stateName,
			"state_code", stateCode,
			"city_name", cityName,
			"city_code", cityCode,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AddLocation(ctx, countryName, countryCode, stateName, stateCode, cityName, cityCode)
}
