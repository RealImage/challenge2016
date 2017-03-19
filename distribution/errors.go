package distribution

import (
	"../errors"
)

func DistributionScopeError(location string) errors.ApplicationError {
	return &errors.BaseApplicationError{
		Code:    "distribution.scope.error",
		Message: "Access denied to location: " + location,
	}
}
