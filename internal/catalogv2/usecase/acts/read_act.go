package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/errors"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/eventbus"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/logger"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/validator"
)

// ReadActInput is the input of the usecase
type ReadActInput struct {
	ID string `json:"id,omitempty" bson:"_id" validate:"required"`
}

// ReadActOutput is the output of the usecase
type ReadActOutput struct {
	*entity.Act
}

// ReadActUC is the use case for getting a musical act by id
type ReadActUC struct {
	ActRepository      repository.ActRepository
	AssetsRepository   repository.AssetsRepository
	RawaudioRepository repository.RawaudioRepository
	Logger             logger.Interface
	Validator          validator.Interface
	Producer           eventbus.Producer
}

// ReadActByIDOpts type is used for implement the command pattern
type ReadActByIDOpts func(*ReadActUC)

// WithGetAcByIDRepository adds the repository to the usecase

// WithGetAcByIDLogger adds the logger to the usecase
func WithGetAcByIDLogger(logg logger.Interface) ReadActByIDOpts {
	return func(uc *ReadActUC) {
		uc.Logger = logg
	}
}

// WithGetAcByIDValidator adds the validator to the usecase
func WithGetAcByIDValidator(val validator.Interface) ReadActByIDOpts {
	return func(uc *ReadActUC) {
		uc.Validator = val
	}
}

// NewReadActByID is the constructor for ReadActUC uscase
func NewReadActByID(opts ...ReadActByIDOpts) *ReadActUC {
	uc := &ReadActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute perform the CreateAct use case.
func (uc *ReadActUC) Execute(ctx context.Context, input *ReadActInput) (*ReadActOutput, error) {
	uc.Logger.Info("Getting a musical act by id")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warnf("invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	act, err := uc.ActRepository.ReadOne(ctx, input.ID)
	if err != nil {
		uc.Logger.Errorf("error getting act: %v", err)
		return nil, err
	}

	uc.Logger.Info("Musical act gotten successfully")
	return &ReadActOutput{act}, nil
}
