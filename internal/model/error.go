package model

import "errors"

var (
	ErrValidation   = errors.New("VALIDATION ERROR")
	ErrNotFound     = errors.New("NOT FOUND")
	ErrInternal     = errors.New("INTERNAL ERROR")
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)
