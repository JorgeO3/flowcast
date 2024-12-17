package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/errors"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/logger"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/validator"
)

// ReadManyInput represents the input parameters for retrieving a list of musical acts.
// It includes pagination parameters: Limit and Offset.
type ReadManyInput struct {
	Limit  int64  `json:"limit,omitempty" bson:"limit"`
	Offset int64  `json:"offset,omitempty" bson:"offset"`
	Genre  string `json:"genre,omitempty" bson:"genre" validate:"omitempty"`
}

// ReadManyOutput encapsulates the list of retrieved musical acts.
type ReadManyOutput struct {
	Acts []*entity.Act
}

// ReadManyUC encapsulates the use case for fetching a list of musical acts.
// It depends on ActRepository for data access, Logger for logging, and Validator for input validation.
type ReadManyUC struct {
	ActRepo      repository.ActRepository
	AssetsRepo   repository.AssetsRepository
	RawaudioRepo repository.RawaudioRepository
	Logger       logger.Interface
	Validator    validator.Interface
}

// ReadManyOpts defines a functional option for configuring ReadManyUC.
// This pattern allows for flexible and readable dependency injection.
type ReadManyOpts func(*ReadManyUC)

// WithReadManyActRepo injects the ActRepository into the use case.
// It allows the use case to fetch musical acts from the repository.
func WithReadManyActRepo(repo repository.ActRepository) ReadManyOpts {
	return func(uc *ReadManyUC) {
		uc.ActRepo = repo
	}
}

// WithReadManyAssetsRepo injects the AssetsRepository into the use case.
// It allows the use case to fetch musical acts from the repository.
func WithReadManyAssetsRepo(repo repository.AssetsRepository) ReadManyOpts {
	return func(uc *ReadManyUC) {
		uc.AssetsRepo = repo
	}
}

// WithReadManyRawaudioRepo injects the RawaudioRepository into the use case.
// It allows the use case to fetch musical acts from the repository.
func WithReadManyRawaudioRepo(repo repository.RawaudioRepository) ReadManyOpts {
	return func(uc *ReadManyUC) {
		uc.RawaudioRepo = repo
	}
}

// WithReadManyLogger injects the Logger into the use case.
// It allows the use case to log informational, warning, and error messages.
func WithReadManyLogger(logg logger.Interface) ReadManyOpts {
	return func(uc *ReadManyUC) {
		uc.Logger = logg
	}
}

// WithReadManyValidator injects the Validator into the use case.
// It ensures that input parameters are validated before processing.
func WithReadManyValidator(val validator.Interface) ReadManyOpts {
	return func(uc *ReadManyUC) {
		uc.Validator = val
	}
}

// NewReadMany creates a new instance of ReadManyUC with the provided functional options.
// This constructor promotes flexibility and decouples the use case from its dependencies,
// making it easier to test and maintain.
func NewReadMany(opts ...ReadManyOpts) *ReadManyUC {
	uc := &ReadManyUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute performs the use case to retrieve a list of musical acts.
// It validates the input, fetches acts from the repository with pagination,
// and returns the result or an appropriate error.
func (uc *ReadManyUC) Execute(ctx context.Context, input ReadManyInput) (*ReadManyOutput, error) {
	uc.Logger.Info("Getting a list of musical acts")

	// Validate input parameters to ensure required fields are present and correct.
	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Errorf("invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Fetch acts from the repository based on the provided limit and offset for pagination.
	acts, err := uc.ActRepo.ReadMany(ctx, input.Genre, input.Limit, input.Offset)
	if err != nil {
		uc.Logger.Errorf("error getting acts: %v", err)
		return nil, err
	}

	if len(acts) == 0 {
		uc.Logger.Warn("no acts found")
		return nil, errors.NewNotFound("no acts found", nil)
	}

	// Return the retrieved acts encapsulated in ReadManyOutput.
	return &ReadManyOutput{acts}, nil
}
