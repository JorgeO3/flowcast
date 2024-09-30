package usecase

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

// UpdateActInput represents the input for the UpdateAct use case.
type UpdateActInput struct {
	e.Act
}

func (uai UpdateActInput) toEntity() *e.Act {
	return e.NewAct(
		e.WithActID(uai.ID),
		e.WithActName(uai.Name),
		e.WithActType(uai.Type),
		e.WithActGenres(uai.Genres),
		e.WithActAlbums(uai.Albums),
		e.WithActMembers(uai.Members),
		e.WithActBiography(uai.Biography),
		e.WithActDisbandDate(uai.DisbandDate),
		e.WithActFormationDate(uai.FormationDate),
		e.WithActProfilePictureURL(uai.ProfilePictureURL),
	)
}

// UpdateActOutput represents the output for the UpdateAct use case.
type UpdateActOutput struct{}

// UpdateActUC is the use case for updating an musical actor.
type UpdateActUC struct {
	ActRepository repository.ActRepository
	Logger        logger.Interface
	Validator     validator.Validator
}

// UpdateActUCOpts represents the functional options for the UpdateActUC.
type UpdateActUCOpts func(uc *UpdateActUC)

// WithUpdateActRepository sets the ActRepository in the UpdateActUC.
func WithUpdateActRepository(repo repository.ActRepository) UpdateActUCOpts {
	return func(uc *UpdateActUC) {
		uc.ActRepository = repo
	}
}

// WithUpdateActLogger sets the logger in the UpdateActUC.
func WithUpdateActLogger(logger logger.Interface) UpdateActUCOpts {
	return func(uc *UpdateActUC) {
		uc.Logger = logger
	}
}

// WithUpdateActValidator sets the validator in the UpdateActUC.
func WithUpdateActValidator(validator validator.Validator) UpdateActUCOpts {
	return func(uc *UpdateActUC) {
		uc.Validator = validator
	}
}

// NewUpdateAct creates a new instance of UpdateActUC.
func NewUpdateAct(opts ...UpdateActUCOpts) *UpdateActUC {
	uc := &UpdateActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute perform the UpdateAct use case.
func (uc *UpdateActUC) Execute(ctx context.Context, input UpdateActInput) (*UpdateActOutput, error) {
	uc.Logger.Info("Updating act in the catalog")

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	if err := uc.ActRepository.UpdateAct(ctx, input.toEntity()); err != nil {
		uc.Logger.Error("Failed to update act: %v", err)
		return nil, errors.HandleRepoError(err)
	}

	uc.Logger.Info("Act updated successfully")
	return nil, nil
}
