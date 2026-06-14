package exception

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation   = errors.New("VALIDATION ERROR")
	ErrNotFound     = errors.New("NOT FOUND")
	ErrInternal     = errors.New("INTERNAL ERROR")
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)

func ExtractValidationErrors(err error) map[string]string {
	errorReport := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errorReport[fieldError.Field()] = fieldError.Tag()
		}
	}

	return errorReport
}

type ClientError interface {
	Code() int
	Error() string
	GetError() any
}

//Error Validation
type ValidationErrorWithFields struct {
	Message string
	Errors  map[string]string
}

func (e *ValidationErrorWithFields) Code() int {
	return 400
}

func (e *ValidationErrorWithFields) Error() string {
	return e.Message
}

func (e *ValidationErrorWithFields) GetError() any {
	return e.Errors
}


//Error Not Found
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Code() int {
	return 404
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *NotFoundError) GetError() any {
	return nil
}


//Error Unauthorized
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Code() int {
	return 401
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) GetError() any {
	return nil
}


// Error Internal
type InternalError struct {
	Message string
}

func (e *InternalError) Code() int {
	return 500
}

func (e *InternalError) Error() string {
	return e.Message
}

func (e *InternalError) GetError() any {
	return nil
}