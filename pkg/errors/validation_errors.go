package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors struct {
	Errors []validationError `json:"errors"`
}

type validationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

func NewValidationErrors(err error) error {
	var errs []validationError

	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, validationError{
			Key:   err.Field(),
			Error: fmt.Sprintf("field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()),
		})
	}

	return &ValidationErrors{
		Errors: errs,
	}
}

func (ve *ValidationErrors) Error() string {
	var strErrors []string
	for _, e := range ve.Errors {
		strErrors = append(strErrors, fmt.Sprintf("%s: %s", e.Key, e.Error))
	}
	return fmt.Sprintf("%v", strErrors)
}
