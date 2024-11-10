// Package errors provides a structure and types for consistent service error handling.
package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrValidation is returned when there are validation errors
	ErrValidation = errors.New("validation error")
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = errors.New("resource not found")
	// ErrConflict is returned when a resource has a conflict
	ErrConflict = errors.New("resource conflict")
	// ErrInternal is returned when an internal server error occurs
	ErrInternal = errors.New("internal server error")
	// ErrUnauthorized is returned when a user is not authorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden is returned when a user is forbidden from accessing a resource
	ErrForbidden = errors.New("forbidden")
	// ErrBadRequest is returned when a request is malformed
	ErrBadRequest = errors.New("bad request")
	// ErrTimeout is returned when a request times out
	ErrTimeout = errors.New("timeout")
)

// DomainError is a custom error type for domain specific errors
type DomainError struct {
	Op  string // Operación realizada
	Err error  // Error original
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

// NewDomainError creates a new domain error
func NewDomainError(op string, err error) *DomainError {
	return &DomainError{
		Op:  op,
		Err: err,
	}
}
