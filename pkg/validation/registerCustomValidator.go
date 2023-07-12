package validation

import (
	"chng2016/pkg/utils"

	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidationFunction ...
func (v *Validation) RegistorCustomValidationFunction() error {
	err := v.Validator.RegisterValidation("isValidStateCode", v.isValidStateCode)
	if err != nil {
		return utils.ErrInValidStateCode
	}

	err = v.Validator.RegisterValidation("isValidCityCode", v.isValidCityCode)
	if err != nil {
		return utils.ErrInValidCityCode
	}

	err = v.Validator.RegisterValidation("isValidCountryCode", v.isValidCountryCode)
	if err != nil {
		return utils.ErrInValidCountryCode
	}

	return nil
}

// isValidStateCode ...
func (v *Validation) isValidStateCode(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return v.utils.IsValidStateCode(value)
}

// isValidCountryCode ...
func (v *Validation) isValidCountryCode(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return v.utils.IsValidCountryCode(value)
}

// isValidCityCode ...
func (v *Validation) isValidCityCode(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return v.utils.IsValidCityCode(value)
}
