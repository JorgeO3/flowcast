package mongodb

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// ErrorCode represents a MongoDB error code.
type ErrorCode string

// MongoDB error codes.
const (
	InternalError                          ErrorCode = "1"
	BadValue                               ErrorCode = "2"
	NoSuchKey                              ErrorCode = "4"
	GraphContainsCycle                     ErrorCode = "5"
	HostUnreachable                        ErrorCode = "6"
	HostNotFound                           ErrorCode = "7"
	UnknownError                           ErrorCode = "8"
	FailedToParse                          ErrorCode = "9"
	CannotMutateObject                     ErrorCode = "10"
	UserNotFound                           ErrorCode = "11"
	UnsupportedFormat                      ErrorCode = "12"
	Unauthorized                           ErrorCode = "13"
	TypeMismatch                           ErrorCode = "14"
	Overflow                               ErrorCode = "15"
	InvalidLength                          ErrorCode = "16"
	ProtocolError                          ErrorCode = "17"
	AuthenticationFailed                   ErrorCode = "18"
	CannotReuseObject                      ErrorCode = "19"
	IllegalOperation                       ErrorCode = "20"
	EmptyArrayOperation                    ErrorCode = "21"
	InvalidBSON                            ErrorCode = "22"
	AlreadyInitialized                     ErrorCode = "23"
	LockTimeout                            ErrorCode = "24"
	RemoteValidationError                  ErrorCode = "25"
	NamespaceNotFound                      ErrorCode = "26"
	IndexNotFound                          ErrorCode = "27"
	PathNotViable                          ErrorCode = "28"
	NonExistentPath                        ErrorCode = "29"
	InvalidPath                            ErrorCode = "30"
	RoleNotFound                           ErrorCode = "31"
	RolesNotRelated                        ErrorCode = "32"
	PrivilegeNotFound                      ErrorCode = "33"
	CannotBackfillArray                    ErrorCode = "34"
	UserModificationFailed                 ErrorCode = "35"
	RemoteChangeDetected                   ErrorCode = "36"
	FileRenameFailed                       ErrorCode = "37"
	FileNotOpen                            ErrorCode = "38"
	FileStreamFailed                       ErrorCode = "39"
	ConflictingUpdateOperators             ErrorCode = "40"
	FileAlreadyOpen                        ErrorCode = "41"
	LogWriteFailed                         ErrorCode = "42"
	CursorNotFound                         ErrorCode = "43"
	UserDataInconsistent                   ErrorCode = "45"
	LockBusy                               ErrorCode = "46"
	NoMatchingDocument                     ErrorCode = "47"
	NamespaceExists                        ErrorCode = "48"
	InvalidRoleModification                ErrorCode = "49"
	MaxTimeMSExpired                       ErrorCode = "50"
	ManualInterventionRequired             ErrorCode = "51"
	DollarPrefixedFieldName                ErrorCode = "52"
	InvalidIDField                         ErrorCode = "53"
	NotSingleValueField                    ErrorCode = "54"
	InvalidDBRef                           ErrorCode = "55"
	EmptyFieldName                         ErrorCode = "56"
	DottedFieldName                        ErrorCode = "57"
	RoleModificationFailed                 ErrorCode = "58"
	CommandNotFound                        ErrorCode = "59"
	ShardKeyNotFound                       ErrorCode = "61"
	OplogOperationUnsupported              ErrorCode = "62"
	StaleShardVersion                      ErrorCode = "63"
	WriteConcernFailed                     ErrorCode = "64"
	MultipleErrorsOccurred                 ErrorCode = "65"
	ImmutableField                         ErrorCode = "66"
	CannotCreateIndex                      ErrorCode = "67"
	IndexAlreadyExists                     ErrorCode = "68"
	AuthSchemaIncompatible                 ErrorCode = "69"
	ShardNotFound                          ErrorCode = "70"
	ReplicaSetNotFound                     ErrorCode = "71"
	InvalidOptions                         ErrorCode = "72"
	InvalidNamespace                       ErrorCode = "73"
	NodeNotFound                           ErrorCode = "74"
	WriteConcernLegacyOK                   ErrorCode = "75"
	NoReplicationEnabled                   ErrorCode = "76"
	OperationIncomplete                    ErrorCode = "77"
	CommandResultSchemaViolation           ErrorCode = "78"
	UnknownReplWriteConcern                ErrorCode = "79"
	RoleDataInconsistent                   ErrorCode = "80"
	NoMatchParseContext                    ErrorCode = "81"
	NoProgressMade                         ErrorCode = "82"
	RemoteResultsUnavailable               ErrorCode = "83"
	IndexOptionsConflict                   ErrorCode = "85"
	IndexKeySpecsConflict                  ErrorCode = "86"
	CannotSplit                            ErrorCode = "87"
	NetworkTimeout                         ErrorCode = "89"
	CallbackCanceled                       ErrorCode = "90"
	ShutdownInProgress                     ErrorCode = "91"
	SecondaryAheadOfPrimary                ErrorCode = "92"
	InvalidReplicaSetConfig                ErrorCode = "93"
	NotYetInitialized                      ErrorCode = "94"
	NotSecondary                           ErrorCode = "95"
	OperationFailed                        ErrorCode = "96"
	NoProjectionFound                      ErrorCode = "97"
	DBPathInUse                            ErrorCode = "98"
	UnsatisfiableWriteConcern              ErrorCode = "100"
	OutdatedClient                         ErrorCode = "101"
	IncompatibleAuditMetadata              ErrorCode = "102"
	NewReplicaSetConfigurationIncompatible ErrorCode = "103"
	NodeNotElectable                       ErrorCode = "104"
	IncompatibleShardingMetadata           ErrorCode = "105"
	DistributedClockSkewed                 ErrorCode = "106"
	LockFailed                             ErrorCode = "107"
	InconsistentReplicaSetNames            ErrorCode = "108"
	ConfigurationInProgress                ErrorCode = "109"
	CannotInitializeNodeWithData           ErrorCode = "110"
	NotExactValueField                     ErrorCode = "111"
	WriteConflict                          ErrorCode = "112"
	InitialSyncFailure                     ErrorCode = "113"
	CommandNotSupported                    ErrorCode = "115"
	ConflictingOperationInProgress         ErrorCode = "117"
	NamespaceNotSharded                    ErrorCode = "118"
	OplogStartMissing                      ErrorCode = "120"
	DocumentValidationFailure              ErrorCode = "121"
	CommandFailed                          ErrorCode = "125"
	CappedPositionLost                     ErrorCode = "136"
	ClientMetadataMissingField             ErrorCode = "183"
	DuplicateKey                           ErrorCode = "11000"
	UnknowedError                          ErrorCode = "UKWN"
)

// Error represents a MongoDB error.
type Error struct {
	Code    ErrorCode
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// NewError creates a new custom mongodb instance of Error.
func NewError(code ErrorCode, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

// MapError maps a MongoDB error to a custom Error.
func MapError(err error) Error {
	if err == nil {
		return Error{}
	}

	var errorMessage strings.Builder
	var errorCode ErrorCode

	switch typedErr := err.(type) {
	case mongo.CommandError:
		errorCode = ErrorCode(typedErr.Code)
		errorMessage.WriteString(fmt.Sprintf("MongoDB command error - Code: %d - Type: %s", typedErr.Code, typedErr.Name))

	case mongo.WriteException:
		if len(typedErr.WriteErrors) > 0 {
			errorCode = ErrorCode(typedErr.WriteErrors[0].Code)
			errorMessage.WriteString(fmt.Sprintf("MongoDB write error - Code: %d - Message: %s",
				typedErr.WriteErrors[0].Code, typedErr.WriteErrors[0].Message))
		}

	case mongo.BulkWriteException:
		if len(typedErr.WriteErrors) > 0 {
			errorCode = ErrorCode(typedErr.WriteErrors[0].Code)
			errorMessage.WriteString(fmt.Sprintf("MongoDB bulk write error - Code: %d - Message: %s",
				typedErr.WriteErrors[0].Code, typedErr.WriteErrors[0].Message))
		}

	case mongo.WriteError:
		errorCode = ErrorCode(typedErr.Code)
		errorMessage.WriteString(fmt.Sprintf("MongoDB write error - Code: %d - Message: %s",
			typedErr.Code, typedErr.Message))

	default:
		errorCode = UnknowedError
		errorMessage.WriteString(fmt.Sprintf("Unknown MongoDB error - Message: %s", err.Error()))
	}

	errorMessage.WriteString(" - Suggestion: Please check your MongoDB operation and try again")

	return NewError(errorCode, errorMessage.String())
}
