package usecase

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteActInput represents the input parameters required to delete a musical act.
// It includes the unique identifier (ID) of the act to be deleted.
type DeleteActInput struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id" validate:"required"`
}

// DeleteActOutput represents the output of the delete act use case.
// Since deleting an act doesn't return any specific data, this struct is empty.
type DeleteActOutput struct{}

// DeleteActUC encapsulates the use case for deleting a musical act.
// It depends on ActRepository for data access, Logger for logging activities,
// and Validator for input validation.
type DeleteActUC struct {
	ActRepository act.Repository
	Logger        logger.Interface
	Validator     validator.Interface
	RaRepository  rawaudio.Repository
	Producer      redpanda.Producer
}

// DeleteActOpts defines a functional option for configuring DeleteActUC.
// This pattern allows for flexible and readable dependency injection.
type DeleteActOpts func(*DeleteActUC)

// WithDeleteActRepository injects the ActRepository into the use case.
// It enables the use case to interact with the data layer for deleting acts.
func WithDeleteActRepository(repo act.Repository) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.ActRepository = repo
	}
}

// WithDeleteActLogger injects the Logger into the use case.
// It allows the use case to log informational, warning, and error messages.
func WithDeleteActLogger(logg logger.Interface) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Logger = logg
	}
}

// WithDeleteActValidator injects the Validator into the use case.
// It ensures that input parameters are validated before processing.
func WithDeleteActValidator(val validator.Interface) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Validator = val
	}
}

// WithDeleteActRaRepository injects the RawAudioRepository into the use case.
// It enables the use case to interact with the raw audio data layer for deleting acts.
func WithDeleteActRaRepository(repo rawaudio.Repository) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.RaRepository = repo
	}
}

// WithDeleteActProducer injects the Producer into the use case.
// It allows the use case to produce messages to a message broker.
func WithDeleteActProducer(prod redpanda.Producer) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Producer = prod
	}
}

// NewDeleteAct creates a new instance of DeleteActUC with the provided functional options.
// This constructor promotes flexibility and decouples the use case from its dependencies,
// making it easier to test and maintain.
func NewDeleteAct(opts ...DeleteActOpts) *DeleteActUC {
	uc := &DeleteActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

// Execute performs the use case to delete a musical act.
// It validates the input, deletes the act from the repository,
// and returns the result or an appropriate error.
func (uc *DeleteActUC) Execute(ctx context.Context, input DeleteActInput) (*DeleteActOutput, error) {
	uc.Logger.Info("Deleting act from the catalog", "id", input.ID.Hex())

	// Validate input parameters to ensure required fields are present and correct.
	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input parameters", "error", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Attempt to delete the act from the repository using the provided ID.
	if err := uc.ActRepository.DeleteAct(ctx, input.ID); err != nil {
		uc.Logger.Error("Failed to delete act from repository", "error", err, "id", input.ID.Hex())
		return nil, errors.HandleRepoError(err)
	}

	// Postear un evento para borrar las canciones del bucket

	// Return an empty output indicating successful deletion.
	return &DeleteActOutput{}, nil
}
