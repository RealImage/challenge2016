package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"chng2016/pkg/utils"

	"github.com/go-playground/validator/v10"
)

// Validator ...
type Validator interface{}

// Validation ...
type Validation struct {
	Validator *validator.Validate
	utils     utils.Util
}

func NewValidation(utils utils.Util) *Validation {
	return &Validation{utils: utils, Validator: validator.New()}
}

var customErrors = map[string]error{
	"DistributorID.required":         errors.New("is required"),
	"Permission.required":            errors.New("is required"),
	"SubDistributorID.required":      errors.New("is required"),
	"CountryCode.required":           errors.New("is required"),
	"StateCode.required":             errors.New("is required"),
	"CityCode.required":              errors.New("is required"),
	"CountryCode.isValidCountryCode": utils.ErrInValidCountryCode,
	"StateCode.isValidStateCode":     utils.ErrInValidStateCode,
	"CityCode.isValidCityCode":       utils.ErrInValidCityCode,
}

// CustomValidationError ...
func (v *Validation) CustomValidationError(sourceStruct interface{}, err error) []map[string]string {
	errs := make([]map[string]string, 0)
	switch err.(type) {
	case validator.ValidationErrors:
		for _, e := range err.(validator.ValidationErrors) {
			errMap := make(map[string]string)
			key := e.Field() + "." + e.Tag()
			fmt.Println("key : ", key)
			if v, ok := customErrors[key]; ok {
				errMap[e.Field()] = v.Error()
			} else {
				errMap[e.Field()] = fmt.Sprintf("custom message is not available : %v", err)
			}
			errs = append(errs, errMap)
		}
		return errs
	case *json.UnmarshalTypeError:
		e := err.(*json.UnmarshalTypeError)
		errs = append(errs, map[string]string{e.Field: fmt.Sprintf("%v can not be a %v", e.Field, e.Value)})
		return errs
	}
	errs = append(errs, map[string]string{"unknown": fmt.Sprintf("unsupported custom error for: %v", err)})
	return errs
}
