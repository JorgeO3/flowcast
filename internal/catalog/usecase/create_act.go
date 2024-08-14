// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"

	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/errors"
	"gitlab.com/JorgeO3/flowcast/internal/catalog/repository"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	ActRepository repository.ActRepository
	Logger        logger.Interface
}

// CreateActUCOpts represents the functional options for the CreateActUC.
type CreateActUCOpts func(uc *CreateActUC)

// WithActRepository sets the ActRepository in the CreateActUC.
func WithActRepository(repo repository.ActRepository) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.ActRepository = repo
	}
}

// WithLogger sets the logger in the CreateActUC.
func WithLogger(logger logger.Interface) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Logger = logger
	}
}

// NewCreateAct creates a new instance of CreateActUC.
func NewCreateAct(actRepo repository.ActRepository, logger logger.Interface) *CreateActUC {
	return &CreateActUC{
		ActRepository: actRepo,
		Logger:        logger,
	}
}

// Execute executes the CreateAct use case.
func (uc *CreateActUC) Execute(ctx context.Context, input CreateActInput) (*CreateActOutput, error) {
	uc.Logger.Info("Executing CreateAct use case")

	if err := validateInput(input); err != nil {
		uc.Logger.Warn("Invalid input data for user registration", "error", err)
		return &CreateActOutput{}, errors.NewValidation("Invalid input data", err)
	}

	act, err := createActEntity(input)
	if err != nil {
		uc.Logger.Error("Failed to create act entity", "error", err)
		return &CreateActOutput{}, errors.NewInternal("Failed to create act entity", err)
	}

	id, err := uc.ActRepository.Create(ctx, act)
	if err != nil {
		uc.Logger.Error("Failed to create act", "error", err)
		return &CreateActOutput{}, errors.NewInternal("Failed to insert the act in the database", err)
	}

	return &CreateActOutput{ID: id}, nil
}

// validateInput validates the input for the CreateAct use case.
func validateInput(input CreateActInput) error {
	_, err := govalidator.ValidateStruct(input)
	return err
}

func createActEntity(input CreateActInput) (*entity.Act, error) {
	return &entity.Act{}, nil
}
