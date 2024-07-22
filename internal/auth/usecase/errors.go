package usecase

import (
	"net/http"

	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
)

type ErrorType int

const (
	ErrorTypeValidation ErrorType = iota
	ErrorTypeNotFound
	ErrorTypeConflict
	ErrorTypeInternal
)

// ErrorType -.
func Hola(err error) ErrorType {
	switch err.(type) {
	case postgres.Error:
		return ErrorTypeConflict
	default:
		return ErrorTypeInternal
	}
}

type DomainError struct {
	Type    ErrorType
	Message string
	Err     error
	Code    int
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

// HttpStatusCode -.
func (e DomainError) HttpStatusCode() int {
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
