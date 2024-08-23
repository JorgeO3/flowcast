// Package usecase provides the use cases for the catalog service.
package usecase

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// CreateActInput represents the input for the CreateAct use case.
type CreateActInput struct {
	e.Act
}

func (cai CreateActInput) toEntity() (*e.Act, error) {
	return e.NewAct(
		e.WithActID(cai.ID),
		e.WithActName(cai.Name),
		e.WithActType(cai.Type),
		e.WithActGenres(cai.Genres),
		e.WithActAlbums(cai.Albums),
		e.WithActMembers(cai.Members),
		e.WithActBiography(cai.Biography),
		e.WithActDisbandDate(cai.DisbandDate),
		e.WithActFormationDate(cai.FormationDate),
		e.WithActProfilePictureURL(cai.ProfilePictureURL),
	)
}

// CreateActOutput represents the output for the CreateAct use case.
type CreateActOutput struct {
	ID string
}

// CreateActUC is the use case for creating an musical actor.
type CreateActUC struct {
	ActRepository repository.ActRepository
	Logger        logger.Interface
	Validator     validator.Validator
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
func WithActLogger(logger logger.Interface) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Logger = logger
	}
}

// WithValidator sets the validator in the CreateActUC.
func WithActValidator(validator validator.Validator) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Validator = validator
	}
}

// NewCreateAct creates a new instance of CreateActUC.
func NewCreateAct(opts ...CreateActUCOpts) *CreateActUC {
	uc := &CreateActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute executes the CreateAct use case.
func (uc *CreateActUC) Execute(ctx context.Context, input CreateActInput) (*CreateActOutput, error) {
	uc.Logger.Info("Executing CreateAct use case")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input data for user registration - %s", err)
		return &CreateActOutput{}, errors.NewValidation("Invalid input data", err)
	}

	act, err := input.toEntity()
	if err != nil {
		uc.Logger.Error("Failed to create act entity - %s", err)
		return &CreateActOutput{}, errors.NewValidation("Failed to create act entity", err)
	}

	id, err := uc.ActRepository.CreateAct(ctx, act)
	if err != nil {
		uc.Logger.Error("Failed to create act error - %s", err)
		return &CreateActOutput{}, errors.NewInternal("Failed to insert the act in the database", err)
	}

	return &CreateActOutput{ID: id}, nil
}
