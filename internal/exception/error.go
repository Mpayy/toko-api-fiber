package exception

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation      = errors.New("Validation Error")
	ErrNotFound        = errors.New("Not Found")
	ErrDuplicatedEmail = errors.New("Email Already Exist")
	ErrInternal        = errors.New("Internal Server Error")
	ErrUnauthorized    = errors.New("Invalid credentials")
)

func ExtractValidationErrors(err error) map[string]string {
	errorReport := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorReport = TranslateValidationError(validationErrors)
	}

	return errorReport
}

func TranslateValidationError(valErr validator.ValidationErrors) map[string]string {
	fieldError := make(map[string]string)
	for _, e := range valErr {
		var message string
		switch e.Tag() {
		case "required":
			message = "must be filled"
		case "email":
			message = "must be a valid email"
		case "min":
			message = "must be at least " + e.Param() + " characters long"
		case "max":
			message = "must be at most " + e.Param() + " characters long"
		case "numeric":
			message = "must be a number"
		case "gt":
			message = "must be greater than " + e.Param()
		case "valid_price":
			message = "price must be multiples of 100"
		default:
			message = "invalid input value"
		}
		fieldError[strings.ToLower(e.Field())] = message
	}

	return fieldError
}

type ClientError interface {
	Code() int
	Error() string
	GetError() any
}

// Error Validation
type ValidationErrorWithFields struct {
	Message string
	Errors  map[string]string
}

func (e *ValidationErrorWithFields) Code() int {
	return 400
}

func (e *ValidationErrorWithFields) Error() string {
	return ErrValidation.Error()
}

func (e *ValidationErrorWithFields) GetError() any {
	return e.Errors
}

func ValidPrice(field validator.FieldLevel) bool {
	value := field.Field().Int()

	if value%100 == 0 {
		return true
	}

	return false
}
