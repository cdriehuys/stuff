package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func validationError(err error) APIError {
	ve := err.(validator.ValidationErrors)

	fieldErrors := make([]FieldError, len(ve))
	for i, fe := range ve {
		fieldErrors[i] = FieldError{Field: fe.Field(), Message: fieldErrorMessage(fe)}
	}

	return APIError{Fields: &fieldErrors}
}

func fieldErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "max":
		return fmt.Sprintf("This field has a maximum length of %s.", fe.Param())

	case "required":
		return "This field is required."
	}

	return ""
}
