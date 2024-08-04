package usecase

import (
	"net/http"
)

type ErrorType int

const (
	ErrorTypeValidation ErrorType = iota
	ErrorTypeNotFound
	ErrorTypeConflict
	ErrorTypeInternal
)

// DomainError -.
type DomainError struct {
	Type    ErrorType
	Message string
	Err     error
}

// NewDomainError -.
func NewDomainError() *DomainError {
	return &DomainError{}
}

// Error -.
func (e DomainError) Error() string {
	return e.Message
}

// Unwrap -.
func (e DomainError) Unwrap() error {
	return e.Err
}

// HTTPStatusCode -.
func (e DomainError) HTTPStatusCode() int {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// HTTPStatusText -.
func (e DomainError) HTTPStatusText() string {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusText(http.StatusBadRequest)
	case ErrorTypeNotFound:
		return http.StatusText(http.StatusNotFound)
	case ErrorTypeConflict:
		return http.StatusText(http.StatusConflict)
	default:
		return http.StatusText(http.StatusInternalServerError)
	}
}
