package http

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/RealImage/challenge2016/location/locationService"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
)

func New(instance string, tracer stdopentracing.Tracer, logger log.Logger) (locationService.Service, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	addLocationEndpoint := httptransport.NewClient(
		http.MethodPost,
		copyURL(u, "/sum"),
		locationService.EncodeHTTPGenericRequest,
		locationService.DecodeAddLocationResponse,
		httptransport.ClientBefore(opentracing.FromHTTPRequest(tracer, "AddLocation", logger)),
	).Endpoint()
	addLocationEndpoint = opentracing.TraceClient(tracer, "AddLocation")(addLocationEndpoint)

	return locationService.Endpoints{
		AddLocationEndpoint: addLocationEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
