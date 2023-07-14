package utils

import "errors"

var (
	// ErrInValidStateCode ...
	ErrInValidStateCode error = errors.New("invalid state code received")
	// ErrInValidCountryCode ...
	ErrInValidCountryCode error = errors.New("invalid country code received")
	// ErrInValidCityCode ...
	ErrInValidCityCode error = errors.New("invalid city code received")
	// ErrInvalidRegion ...
	ErrInvalidRegion error = errors.New("invalid region received")
)
