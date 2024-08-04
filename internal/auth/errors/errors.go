// Package errors provides a structure and types for consistent service error handling.
package errors

import (
	"fmt"
	"net/http"
)

// AuthError -.
type AuthError interface {
	Error() string
	Code() int
	Message() string
	Unwrap() error
}

// baseError is a struct that will be embedded in all our specific error types.
type baseError struct {
	err      string
	original error
}

// Error returns a string representation of the error, including the original error if present.
func (e baseError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s - error: %v", e.err, e.original)
	}
	return e.err
}

// Unwrap returns the original error.
func (e baseError) Unwrap() error {
	return e.original
}

// Validation represents a validation error in the input data.
type Validation struct {
	baseError
}

// NewValidation creates a new Validation error.
func NewValidation(err string, original error) Validation {
	return Validation{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the validation error.
func (e Validation) Code() int { return http.StatusBadRequest }

// Message returns a human-readable error message for validation errors.
func (e Validation) Message() string { return "Validation Error" }

// NotFound represents an error where the requested resource was not found.
type NotFound struct {
	baseError
}

// NewNotFound creates a new NotFound error.
func NewNotFound(err string, original error) AuthError {
	return NotFound{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the not found error.
func (e NotFound) Code() int { return http.StatusNotFound }

// Message returns a human-readable error message for not found errors.
func (e NotFound) Message() string { return "Not Found" }

// Conflict represents a conflict error, such as a uniqueness violation.
type Conflict struct {
	baseError
}

// NewConflict creates a new Conflict error.
func NewConflict(err string, original error) AuthError {
	return Conflict{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the conflict error.
func (e Conflict) Code() int { return http.StatusConflict }

// Message returns a human-readable error message for conflict errors.
func (e Conflict) Message() string { return "Conflict" }

// Internal represents an internal server error.
type Internal struct {
	baseError
}

// NewInternal creates a new Internal error.
func NewInternal(err string, original error) AuthError {
	return Internal{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the internal server error.
func (e Internal) Code() int { return http.StatusInternalServerError }

// Message returns a human-readable error message for internal server errors.
func (e Internal) Message() string { return "Internal Server Error" }

// Unauthorized indicates that the request lacks valid authentication credentials.
type Unauthorized struct {
	baseError
}

// NewUnauthorized creates a new Unauthorized error.
func NewUnauthorized(err string, original error) AuthError {
	return Unauthorized{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the unauthorized error.
func (e Unauthorized) Code() int { return http.StatusUnauthorized }

// Message returns a human-readable error message for unauthorized errors.
func (e Unauthorized) Message() string { return "Unauthorized" }

// Forbidden indicates that the client doesn't have permission to access the requested resource.
type Forbidden struct {
	baseError
}

// NewForbidden creates a new Forbidden error.
func NewForbidden(err string, original error) AuthError {
	return Forbidden{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the forbidden error.
func (e Forbidden) Code() int { return http.StatusForbidden }

// Message returns a human-readable error message for forbidden errors.
func (e Forbidden) Message() string { return "Forbidden" }

// BadRequest indicates a malformed or invalid client request.
type BadRequest struct {
	baseError
}

// NewBadRequest creates a new BadRequest error.
func NewBadRequest(err string, original error) *BadRequest {
	return &BadRequest{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the bad request error.
func (e BadRequest) Code() int { return http.StatusBadRequest }

// Message returns a human-readable error message for bad request errors.
func (e BadRequest) Message() string { return "Bad Request" }

// Timeout indicates that the operation did not complete within the expected time.
type Timeout struct {
	baseError
}

// NewTimeout creates a new Timeout error.
func NewTimeout(err string, original error) AuthError {
	return Timeout{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the timeout error.
func (e Timeout) Code() int { return http.StatusGatewayTimeout }

// Message returns a human-readable error message for timeout errors.
func (e Timeout) Message() string { return "Timeout" }

// TooManyRequests indicates that the client has sent too many requests in a given amount of time.
type TooManyRequests struct {
	baseError
}

// NewTooManyRequests creates a new TooManyRequests error.
func NewTooManyRequests(err string, original error) AuthError {
	return TooManyRequests{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the too many requests error.
func (e TooManyRequests) Code() int { return http.StatusTooManyRequests }

// Message returns a human-readable error message for too many requests errors.
func (e TooManyRequests) Message() string { return "Too Many Requests" }
