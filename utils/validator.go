package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Errors []ErrorValidatorField `json:"errors"`
	}

	ErrorValidatorField struct {
		FailedField string      `json:"failedField"`
		Tag         string      `json:"tag"`
		Value       interface{} `json:"value"`
		Message     string      `json:"message"`
	}

	XValidator struct {
		validator *validator.Validate
	}
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

var Validator = &XValidator{
	validator: validate,
}

func (v XValidator) Validate(data interface{}) ErrorResponse {
	validationErrors := []ErrorValidatorField{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorValidatorField

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Message = fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'",
				elem.FailedField,
				elem.Value,
				elem.Tag)

			validationErrors = append(validationErrors, elem)
		}
	}

	errorResponse := ErrorResponse{validationErrors}

	return errorResponse
}

func NewErrorResponse(err string) ErrorResponse {
	var elem ErrorValidatorField
	elem.Message = err
	fmt.Printf(err)

	return ErrorResponse{[]ErrorValidatorField{elem}}
}
