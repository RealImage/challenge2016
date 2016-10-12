package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RealImage/challenge2016/location/distributionService"
	"github.com/RealImage/challenge2016/location/locationService"
	"github.com/RealImage/challenge2016/location/repository/inmemory"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var httpPort string
var logging bool

func init() {
	flag.StringVar(&httpPort, "p", "8080", "Location Service http port(:port)")
	flag.BoolVar(&logging, "l", false, "Is logging enable.")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	var logger log.Logger
	if logging {
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	} else {
		logger = log.NewNopLogger()
	}

	locationRepo := repository.NewLocationRepository()
	distributorRepo := repository.NewDistributorRepository()

	fieldKeys := []string{"method"}

	var locationSvc locationService.Service
	locationSvc = locationService.NewService(locationRepo)
	locationSvc = locationService.NewLoggingService(log.NewContext(logger).With("component", "locationService"), locationSvc)
	locationSvc = locationService.NewInstrumentingService(
		kitprometheus.NewCounter(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "location_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		metrics.NewTimeHistogram(time.Microsecond, kitprometheus.NewSummary(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "location_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)), locationSvc)

	var distributorSvc distributionService.Service
	distributorSvc = distributionService.NewService(distributorRepo, locationRepo)
	distributorSvc = distributionService.NewLoggingService(log.NewContext(logger).With("component", "distributionService"), distributorSvc)
	distributorSvc = distributionService.NewInstrumentingService(
		kitprometheus.NewCounter(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "distributor_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		metrics.NewTimeHistogram(time.Microsecond, kitprometheus.NewSummary(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "distributor_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys)), distributorSvc)

	httpLogger := log.NewContext(logger).With("component", "http")

	mux := http.NewServeMux()

	mux.Handle("/api/location/", locationService.MakeHandler(ctx, locationSvc, httpLogger))
	mux.Handle("/api/distributor/", distributionService.MakeHandler(ctx, distributorSvc, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", stdprometheus.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", ":"+httpPort, "msg", "listening")
		errs <- http.ListenAndServe(":"+httpPort, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}