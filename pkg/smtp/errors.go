package smtp

import (
	"fmt"
)

// ErrorCode represents an SMTP error code.
type ErrorCode string

// SMTP error codes.
const (
	ErrCodeAuthFailed         ErrorCode = "AUTH_FAILED"
	ErrCodeSendFailed         ErrorCode = "SEND_FAILED"
	ErrCodeRecipientRejected  ErrorCode = "RECIPIENT_REJECTED"
	ErrCodeMailBoxUnavailable ErrorCode = "MAILBOX_UNAVAILABLE"
	ErrCodeInvalidRecipient   ErrorCode = "INVALID_RECIPIENT"
	ErrCodeTransactionFailed  ErrorCode = "TRANSACTION_FAILED"
	ErrUnknownError           ErrorCode = "UNKNOWN"
)

// Error is a custom error type that represents an SMTP error.
type Error struct {
	Code    ErrorCode
	Message string
	Detail  string
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s, detail: %s", e.Code, e.Message, e.Detail)
}

// MapError maps an SMTP error to an SMTPError.
func MapError(err error) Error {
	if err == nil {
		return Error{}
	}

	errMsg := err.Error()

	switch {
	case errMsg == "535 5.7.8 Authentication credentials invalid":
		return Error{
			Code:    ErrCodeAuthFailed,
			Message: "authentication credentials invalid",
			Detail:  errMsg,
		}
	case errMsg == "550 5.1.1 User unknown":
		return Error{
			Code:    ErrCodeSendFailed,
			Message: "user unknown",
			Detail:  errMsg,
		}
	case errMsg == "552 5.2.2 Over quota":
		return Error{
			Code:    ErrCodeMailBoxUnavailable,
			Message: "over quota",
			Detail:  errMsg,
		}
	case errMsg == "553 5.1.8 Sender address rejected":
		return Error{
			Code:    ErrCodeInvalidRecipient,
			Message: "sender address rejected",
			Detail:  errMsg,
		}
	case errMsg == "554 5.7.1 Transaction failed":
		return Error{
			Code:    ErrCodeTransactionFailed,
			Message: "transaction failed",
			Detail:  errMsg,
		}
	default:
		return Error{
			Code:   ErrUnknownError,
			Detail: errMsg,
		}
	}
}
