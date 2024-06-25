package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrorCode represents a PostgreSQL error code.
type ErrorCode string

// PostgreSQL error codes.
const (
	ErrCodeUniqueViolation           ErrorCode = "23505"
	ErrCodeForeignKeyViolation       ErrorCode = "23503"
	ErrCodeNotNullViolation          ErrorCode = "23502"
	ErrCodeCheckViolation            ErrorCode = "23514"
	ErrCodeExclusionViolation        ErrorCode = "23P01"
	ErrCodeSyntaxError               ErrorCode = "42601"
	ErrCodeUndefinedTable            ErrorCode = "42P01"
	ErrCodeInvalidTextRepresentation ErrorCode = "22P02"
	ErrCodeDivisionByZero            ErrorCode = "22012"
	ErrCodeDataTypeMismatch          ErrorCode = "42804"
	ErrCodeInvalidForeignKey         ErrorCode = "42830"
	ErrCodeSerializationFailure      ErrorCode = "40001"
	ErrCodeDeadlockDetected          ErrorCode = "40P01"
	ErrCodeNotFound                  ErrorCode = "P0002" // Custom error code for not found errors
)

// Error is a custom error type that represents a PostgreSQL error.
type Error struct {
	Code    ErrorCode
	Message string
	Detail  string
	Hint    string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s, detail: %s, hint: %s", e.Code, e.Message, e.Detail, e.Hint)
}

// NewPostgresError crea una nueva instancia de PostgresError a partir de un pgconn.PgError.
func NewPostgresError(pgErr *pgconn.PgError) Error {
	return Error{
		Code:    ErrorCode(pgErr.Code),
		Message: pgErr.Message,
		Detail:  pgErr.Detail,
		Hint:    pgErr.Hint,
	}
}

// MapError maps a PostgreSQL error to a PostgresError.
func MapError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return NewPostgresError(pgErr)
	}

	if err == pgx.ErrNoRows {
		return Error{
			Code:    ErrCodeNotFound,
			Message: "not found",
		}
	}

	return err
}
