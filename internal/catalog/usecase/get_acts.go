package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// GetActsInput represents the input parameters for retrieving a list of musical acts.
// It includes pagination parameters: Limit and Offset.
type GetActsInput struct {
	Limit  int64  `json:"limit,omitempty" bson:"limit"`
	Offset int64  `json:"offset,omitempty" bson:"offset"`
	Genre  string `json:"genre,omitempty" bson:"genre" validate:"omitempty"`
}

// GetActsOutput encapsulates the list of retrieved musical acts.
type GetActsOutput struct {
	Acts []*entity.Act
}

// GetActsUC encapsulates the use case for fetching a list of musical acts.
// It depends on ActRepository for data access, Logger for logging, and Validator for input validation.
type GetActsUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
}

// GetActsOpts defines a functional option for configuring GetActsUC.
// This pattern allows for flexible and readable dependency injection.
type GetActsOpts func(*GetActsUC)

// WithGetActsRepository injects the ActRepository into the use case.
// It enables the use case to interact with the data layer for retrieving acts.
func WithGetActsRepository(repo act.Repository) GetActsOpts {
	return func(uc *GetActsUC) {
		uc.ActRepository = repo
	}
}

// WithGetActsLogger injects the Logger into the use case.
// It allows the use case to log informational, warning, and error messages.
func WithGetActsLogger(logg logger.Interface) GetActsOpts {
	return func(uc *GetActsUC) {
		uc.Logger = logg
	}
}

// WithGetActsValidator injects the Validator into the use case.
// It ensures that input parameters are validated before processing.
func WithGetActsValidator(val validator.Interface) GetActsOpts {
	return func(uc *GetActsUC) {
		uc.Validator = val
	}
}

// NewGetActs creates a new instance of GetActsUC with the provided functional options.
// This constructor promotes flexibility and decouples the use case from its dependencies,
// making it easier to test and maintain.
func NewGetActs(opts ...GetActsOpts) *GetActsUC {
	uc := &GetActsUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute performs the use case to retrieve a list of musical acts.
// It validates the input, fetches acts from the repository with pagination,
// and returns the result or an appropriate error.
func (uc *GetActsUC) Execute(ctx context.Context, input GetActsInput) (*GetActsOutput, error) {
	uc.Logger.Info("Fetching all musical acts", "limit", input.Limit, "offset", input.Offset)

	// Validate input parameters to ensure required fields are present and correct.
	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Error("Invalid input parameters", "error", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Fetch acts from the repository based on the provided limit and offset for pagination.
	acts, err := uc.ActRepository.GetActs(ctx, input.Genre, input.Limit, input.Offset)
	if err != nil {
		uc.Logger.Error("Failed to retrieve acts from repository", "error", err)
		return nil, errors.HandleRepoError(err)
	}

	if len(acts) == 0 {
		uc.Logger.Info("No acts found for the given parameters - limit: %d, offset: %d", input.Limit, input.Offset)
		return nil, errors.NewNotFound("no acts found", nil)
	}

	// Return the retrieved acts encapsulated in GetActsOutput.
	return &GetActsOutput{acts}, nil
}
