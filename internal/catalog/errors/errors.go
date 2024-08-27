// Package errors provides a structure and types for consistent service error handling.
package errors

import (
	"fmt"
	"net/http"

	"github.com/JorgeO3/flowcast/pkg/mongodb"
)

// CatalogError -.
type CatalogError interface {
	Error() string
	Code() int
	Msg() string
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

// Msg returns a human-readable error message for validation errors.
func (e Validation) Msg() string { return "Validation Error" }

// NotFound represents an error where the requested resource was not found.
type NotFound struct {
	baseError
}

// NewNotFound creates a new NotFound error.
func NewNotFound(err string, original error) CatalogError {
	return NotFound{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the not found error.
func (e NotFound) Code() int { return http.StatusNotFound }

// Msg returns a human-readable error message for not found errors.
func (e NotFound) Msg() string { return "Not Found" }

// Conflict represents a conflict error, such as a uniqueness violation.
type Conflict struct {
	baseError
}

// NewConflict creates a new Conflict error.
func NewConflict(err string, original error) CatalogError {
	return Conflict{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the conflict error.
func (e Conflict) Code() int { return http.StatusConflict }

// Msg returns a human-readable error message for conflict errors.
func (e Conflict) Msg() string { return "Conflict" }

// Internal represents an internal server error.
type Internal struct {
	baseError
}

// NewInternal creates a new Internal error.
func NewInternal(err string, original error) CatalogError {
	return Internal{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the internal server error.
func (e Internal) Code() int { return http.StatusInternalServerError }

// Msg returns a human-readable error message for internal server errors.
func (e Internal) Msg() string { return "Internal Server Error" }

// Unauthorized indicates that the request lacks valid authentication credentials.
type Unauthorized struct {
	baseError
}

// NewUnauthorized creates a new Unauthorized error.
func NewUnauthorized(err string, original error) CatalogError {
	return Unauthorized{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the unauthorized error.
func (e Unauthorized) Code() int { return http.StatusUnauthorized }

// Msg returns a human-readable error message for unauthorized errors.
func (e Unauthorized) Msg() string { return "Unauthorized" }

// Forbidden indicates that the client doesn't have permission to access the requested resource.
type Forbidden struct {
	baseError
}

// NewForbidden creates a new Forbidden error.
func NewForbidden(err string, original error) CatalogError {
	return Forbidden{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the forbidden error.
func (e Forbidden) Code() int { return http.StatusForbidden }

// Msg returns a human-readable error message for forbidden errors.
func (e Forbidden) Msg() string { return "Forbidden" }

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

// Msg returns a human-readable error message for bad request errors.
func (e BadRequest) Msg() string { return "Bad Request" }

// Timeout indicates that the operation did not complete within the expected time.
type Timeout struct {
	baseError
}

// NewTimeout creates a new Timeout error.
func NewTimeout(err string, original error) CatalogError {
	return Timeout{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the timeout error.
func (e Timeout) Code() int { return http.StatusGatewayTimeout }

// Msg returns a human-readable error message for timeout errors.
func (e Timeout) Msg() string { return "Timeout" }

// TooManyRequests indicates that the client has sent too many requests in a given amount of time.
type TooManyRequests struct {
	baseError
}

// NewTooManyRequests creates a new TooManyRequests error.
func NewTooManyRequests(err string, original error) CatalogError {
	return TooManyRequests{baseError{err: err, original: original}}
}

// Code returns the HTTP status code corresponding to the too many requests error.
func (e TooManyRequests) Code() int { return http.StatusTooManyRequests }

// Msg returns a human-readable error message for too many requests errors.
func (e TooManyRequests) Msg() string { return "Too Many Requests" }

// FIXME: This is a temporary solution until we have a better error handling strategy.
// Check this
// func mapMongoErrorToHTTPStatus(err error) int {
//     if mongoErr, ok := err.(mongo.WriteException); ok {
//         for _, we := range mongoErr.WriteErrors {
//             switch we.Code {
//             case 11000:
//                 return http.StatusConflict // 409
//             case 50, 51:
//                 return http.StatusGatewayTimeout // 504
//             case 13:
//                 return http.StatusForbidden // 403
//             default:
//                 if we.Code >= 0 && we.Code <= 49 {
//                     return http.StatusBadRequest // 400
//                 } else if we.Code >= 10000 && we.Code < 11000 {
//                     return http.StatusInternalServerError // 500
//                 }
//             }
//         }
//     }
//     return http.StatusInternalServerError // 500 por defecto
// }

// HandleRepoError returns a CatalogError based on the error received from the repository.
func HandleRepoError(e error) error {
	if err, ok := e.(mongodb.Error); ok {
		switch err.Code {
		case mongodb.DuplicateKey:
			return NewConflict("Duplicate key", err)
		case mongodb.InternalError:
			return NewInternal("Internal error", err)
		case mongodb.BadValue:
			return NewBadRequest("Bad request", err)
		case mongodb.NoSuchKey:
			return NewNotFound("Not found", err)
		case mongodb.GraphContainsCycle:
			return NewInternal("Graph contains cycle", err)
		case mongodb.HostUnreachable:
			return NewInternal("Host unreachable", err)
		case mongodb.HostNotFound:
			return NewInternal("Host not found", err)
		case mongodb.UnknownError:
			return NewInternal("Unknown error", err)
		case mongodb.FailedToParse:
			return NewInternal("Failed to parse", err)
		case mongodb.CannotMutateObject:
			return NewInternal("Cannot mutate object", err)
		case mongodb.UserNotFound:
			return NewInternal("User not found", err)
		case mongodb.UnsupportedFormat:
			return NewInternal("Unsupported format", err)
		case mongodb.Unauthorized:
			return NewUnauthorized("Unauthorized", err)
		case mongodb.TypeMismatch:
			return NewInternal("Type mismatch", err)
		case mongodb.Overflow:
			return NewInternal("Overflow", err)
		case mongodb.InvalidLength:
			return NewInternal("Invalid length", err)
		case mongodb.ProtocolError:
			return NewInternal("Protocol error", err)
		case mongodb.AuthenticationFailed:
			return NewInternal("Authentication failed", err)
		case mongodb.CannotReuseObject:
			return NewInternal("Cannot reuse object", err)
		case mongodb.IllegalOperation:
			return NewInternal("Illegal operation", err)
		case mongodb.EmptyArrayOperation:
			return NewInternal("Empty array operation", err)
		case mongodb.InvalidBSON:
			return NewInternal("Invalid BSON", err)
		case mongodb.AlreadyInitialized:
			return NewInternal("Already initialized", err)
		case mongodb.LockTimeout:
			return NewInternal("Lock timeout", err)
		case mongodb.RemoteValidationError:
			return NewInternal("Remote validation error", err)
		case mongodb.NamespaceNotFound:
			return NewInternal("Namespace not found", err)
		case mongodb.IndexNotFound:
			return NewInternal("Index not found", err)
		case mongodb.PathNotViable:
			return NewInternal("Path not viable", err)
		default:
			return NewInternal("Internal error", err)
		}
	}

	return e
}
