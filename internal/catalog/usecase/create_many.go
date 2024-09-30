package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateManyInput represents the input for the CreateMany use case.
type CreateManyInput struct {
	Acts []*entity.Act `json:"acts" validate:"required,dive,required"`
}

// CreateManyOutput represents the output for the CreateMany use case.
type CreateManyOutput struct {
	IDs []string `json:"ids"`
}

// CreateManyUC is the use case for creating multiple musical actors.
type CreateManyUC struct {
	ActRepository repository.ActRepository
	Logger        logger.Interface
	Validator     validator.Validator
}

// CreateManyUCOpts represents the functional options for the CreateManyUC.
type CreateManyUCOpts func(uc *CreateManyUC)

// WithCreateManyRepository sets the ActRepository in the CreateManyUC.
func WithCreateManyRepository(repo repository.ActRepository) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.ActRepository = repo
	}
}

// WithCreateManyLogger sets the logger in the CreateManyUC.
func WithCreateManyLogger(logger logger.Interface) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.Logger = logger
	}
}

// WithCreateManyValidator sets the validator in the CreateManyUC.
func WithCreateManyValidator(validator validator.Validator) CreateManyUCOpts {
	return func(uc *CreateManyUC) {
		uc.Validator = validator
	}
}

// NewCreateMany is the constructor for CreateManyUC use case
func NewCreateMany(opts ...CreateManyUCOpts) *CreateManyUC {
	uc := &CreateManyUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateMany use case
func (uc *CreateManyUC) Execute(ctx context.Context, input CreateManyInput) (*CreateManyOutput, error) {
	uc.Logger.Info("Creating multiple musical acts")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	ids, err := uc.ActRepository.CreateManyActs(ctx, input.Acts)
	if err != nil {
		uc.Logger.Error("Error inserting acts in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	return &CreateManyOutput{IDs: ids}, nil
}
