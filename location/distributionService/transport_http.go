package distributionService

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/RealImage/challenge2016/location/domain"
	"github.com/google/go-querystring/query"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"golang.org/x/net/context"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(ctx context.Context, s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(errorEncoder),
	}

	addDistributorHandler := kithttp.NewServer(
		ctx,
		makeAddDistributorEndpoint(s),
		decodeAddDistributorRequest,
		encodeResponse,
		opts...,
	)

	checkLocationPermissionHandler := kithttp.NewServer(
		ctx,
		makeCheckLocationPermissionEndpoint(s),
		decodeCheckLocationPermissionRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/distributor/v1", addDistributorHandler).Methods(http.MethodPost)
	r.Handle("/api/distributor/v1/permission", checkLocationPermissionHandler).Methods(http.MethodGet)

	return r
}

func decodeAddDistributorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addDistributorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil

}

func DecodeAddDistributorResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp addDistributorResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeCheckLocationPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req checkLocationPermissionRequest
	dec := schema.NewDecoder()
	if err := dec.Decode(&req, r.URL.Query()); err != nil {
		return nil, err
	}
	return req, nil

}

func DecodeCheckLocationPermissionResponse(_ context.Context, r *http.Response) (interface{}, error) {

	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp checkLocationPermissionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)

	return resp, err
}

func EncodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func EncodeHTTPGenericGetRequest(_ context.Context, r *http.Request, request interface{}) error {
	val, err := query.Values(request)
	if err != nil {
		return err
	}
	r.URL.RawQuery = val.Encode()
	return nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		errorEncoder(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	msg := err.Error()

	if e, ok := err.(kithttp.Error); ok {
		msg = e.Err.Error()
		switch e.Domain {
		case kithttp.DomainDecode:
			code = http.StatusBadRequest

		case kithttp.DomainDo:
			switch e.Err {
			case domain.ErrInvalidArgument:
				fallthrough
			case domain.ErrParentHaveNotPermission:
				fallthrough
			case domain.ErrInvalidLocation:
				fallthrough
			case domain.ErrDistributorNotFound:
				fallthrough
			case domain.ErrAlreadyHavePermission:
				code = http.StatusBadRequest
				break
			case domain.ErrExists:
				fallthrough
			case domain.ErrNotFound:
				code = http.StatusInternalServerError
				break
			}
		}
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorWrapper{Error: msg})
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
