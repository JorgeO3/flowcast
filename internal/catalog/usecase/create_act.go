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

func (cai CreateActInput) toEntity() *e.Act {
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
func WithCreateActRepository(repo repository.ActRepository) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.ActRepository = repo
	}
}

// WithActLogger sets the logger in the CreateActUC.
func WithCreateActLogger(logger logger.Interface) CreateActUCOpts {
	return func(uc *CreateActUC) {
		uc.Logger = logger
	}
}

// WithActValidator sets the validator in the CreateActUC.
func WithCreateActValidator(validator validator.Validator) CreateActUCOpts {
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
	uc.Logger.Info("Creating a new musical act")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	id, err := uc.ActRepository.CreateAct(ctx, input.toEntity())
	if err != nil {
		uc.Logger.Error("Error inserting act in db: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	uc.Logger.Info("Act created successfully")
	return &CreateActOutput{ID: id}, nil
}
