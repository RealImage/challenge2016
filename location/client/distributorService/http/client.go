package http

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/RealImage/challenge2016/location/distributionService"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
)

func New(instance string, tracer stdopentracing.Tracer, logger log.Logger) (distributionService.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	addDistributorEndpoint := httptransport.NewClient(
		http.MethodPost,
		copyURL(u, "/api/v1/distributor"),
		distributionService.EncodeHTTPGenericRequest,
		distributionService.DecodeAddDistributorResponse,
		httptransport.ClientBefore(opentracing.FromHTTPRequest(tracer, "AddDistributor", logger)),
	).Endpoint()
	addDistributorEndpoint = opentracing.TraceClient(tracer, "AddDistributor")(addDistributorEndpoint)

	return distributionService.Endpoints{
		AddDistributorEndpoint: addDistributorEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
