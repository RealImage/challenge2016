package distributionService

import (
	"time"

	"golang.org/x/net/context"

	"github.com/RealImage/challenge2016/location/domain"
	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.TimeHistogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(requestCount metrics.Counter, requestLatency metrics.TimeHistogram, s Service) Service {
	return &instrumentingService{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		Service:        s,
	}
}

func (s *instrumentingService) AddDistributor(ctx context.Context, parentDistributorId domain.DistributorId, distributorId domain.DistributorId, locationType domain.LocationType, permission domain.Permission, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddDistributor"}
		s.requestCount.With(methodField).Add(1)
		s.requestLatency.With(methodField).Observe(time.Since(begin))
	}(time.Now())

	return s.Service.AddDistributor(ctx, parentDistributorId, distributorId, locationType, permission, countryCode, stateCode, cityCode)
}

func (s *instrumentingService) CheckLocationPermission(ctx context.Context, distributorId domain.DistributorId, countryCode domain.CountryCode, stateCode domain.StateCode, cityCode domain.CityCode) (ok bool, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "CheckLocationPermission"}
		s.requestCount.With(methodField).Add(1)
		s.requestLatency.With(methodField).Observe(time.Since(begin))
	}(time.Now())

	return s.Service.CheckLocationPermission(ctx, distributorId, countryCode, stateCode, cityCode)
}
