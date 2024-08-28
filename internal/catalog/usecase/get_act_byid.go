package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetActByIDInput is the input of the usecase
type GetActByIDInput struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id" validate:"required"`
}

// GetActByIDOutput is the output of the usecase
type GetActByIDOutput struct {
	*entity.Act
}

// GetActByIDUC is the use case for getting a musical act by id
type GetActByIDUC struct {
	ActRepository repository.ActRepository
	Logger        logger.Interface
	Validator     validator.Validator
}

// GetActByIDOpts type is used for implement the command pattern
type GetActByIDOpts func(*GetActByIDUC)

// WithGetAcByIDRepository adds the repo to the usecase
func WithGetAcByIDRepository(repo repository.ActRepository) GetActByIDOpts {
	return func(uc *GetActByIDUC) {
		uc.ActRepository = repo
	}
}

// WithGetAcByIDLogger adds the logger to the usecase
func WithGetAcByIDLogger(logg logger.Interface) GetActByIDOpts {
	return func(uc *GetActByIDUC) {
		uc.Logger = logg
	}
}

// WithGetAcByIDValidator adds the validator to the usecase
func WithGetAcByIDValidator(val validator.Validator) GetActByIDOpts {
	return func(uc *GetActByIDUC) {
		uc.Validator = val
	}
}

// NewGetActByIDUC is the constructor for GetActByIDUC uscase
func NewGetActByID(opts ...GetActByIDOpts) *GetActByIDUC {
	uc := &GetActByIDUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute perform the CreateAct use case.
func (uc *GetActByIDUC) Execute(ctx context.Context, input GetActByIDInput) (*GetActByIDOutput, error) {
	uc.Logger.Info("Getting a musical act by id")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	act, err := uc.ActRepository.GetActByID(ctx, input.ID)
	if err != nil {
		uc.Logger.Error("Failed to get act: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	uc.Logger.Info("Musical act gotten successfully")
	return &GetActByIDOutput{act}, nil
}
