package handlers

import (
	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

var validate = validator.New()

func ValidateStruct(data any) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var res ValidationErrorResponse

			res.FailedField = err.Field()
			res.Tag = err.Tag()
			res.Value = err.Value()
			res.Error = true

			validationErrors = append(validationErrors, res)
		}
	}

	return validationErrors
}
