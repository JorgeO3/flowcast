package usecase

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/mongotx"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/google/uuid"
)

// DeleteActInput represents the input parameters required to delete a musical act.
// It includes the unique identifier (ID) of the act to be deleted.
type DeleteActInput struct {
	ID string `json:"id,omitempty" bson:"_id" validate:"required"`
}

// DeleteActOutput represents the output of the delete act use case.
// Since deleting an act doesn't return any specific data, this struct is empty.
type DeleteActOutput struct{}

// DeleteActEvent represents the event to be posted when an act is deleted.
type DeleteActEvent struct {
	EventID string   `json:"eventId"`
	Assets  []string `json:"assets"`
}

// DeleteActUC encapsulates the use case for deleting a musical act.
// It depends on ActRepository for data access, Logger for logging activities,
// and Validator for input validation.
type DeleteActUC struct {
	Logger    logger.Interface
	Validator validator.Interface
	Producer  redpanda.Producer
	Repos     *repository.Repositories
	tx        *mongotx.MongoTx
}

// DeleteActOpts defines a functional option for configuring DeleteActUC.
// This pattern allows for flexible and readable dependency injection.
type DeleteActOpts func(*DeleteActUC)

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

// WithDeleteActProducer injects the Producer into the use case.
// It allows the use case to produce messages to a message broker.
func WithDeleteActProducer(prod redpanda.Producer) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Producer = prod
	}
}

// WithDeleteActRepositories injects the Repositories into the use case.
// It provides access to the data repositories needed by the use case.
func WithDeleteActRepositories(repos *repository.Repositories) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Repos = repos
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

func (uc *DeleteActUC) handleDelete(ctx context.Context, id string) (*e.Act, error) {

	// Get the act from the repository using the provided ID.
	act, err := uc.Repos.Act.GetActByID(ctx, id)
	if err != nil {
		uc.Logger.Error("Failed to get act from repository", "error", err, "id", id)
		return nil, errors.HandleRepoError(err)
	}

	// Attempt to delete the act from the repository using the provided ID.
	if err := uc.Repos.Act.DeleteAct(ctx, id); err != nil {
		uc.Logger.Error("Failed to delete act from repository", "error", err, "id", id)
		return nil, errors.HandleRepoError(err)
	}

	return act, nil
}

func getAssetIDs(assets []AudioServiceAsset) []string {
	assetIDs := make([]string, len(assets))
	for i, asset := range assets {
		assetIDs[i] = asset.AssetID
	}
	return assetIDs
}

// Execute performs the use case to delete a musical act.
// It validates the input, deletes the act from the repository,
// and returns the result or an appropriate error.
func (uc *DeleteActUC) Execute(ctx context.Context, input DeleteActInput) (*DeleteActOutput, error) {
	uc.Logger.Info("Deleting act from the catalog", "id", input.ID)

	// Validate input parameters to ensure required fields are present and correct.
	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warn("Invalid input parameters", "error", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	var act *e.Act
	var err error
	err = uc.tx.Run(ctx, func(ctx context.Context) error {
		act, err = uc.handleDelete(ctx, input.ID)
		return err
	})

	if err != nil {
		uc.Logger.Error("Failed to start transaction", "error", err)
		return nil, errors.NewInternal("error starting transaction", err)
	}

	processor := utils.NewAssetsProcessor(ctx, uc.Repos)
	output, err := processor.Delete(act)
	if err != nil {
		uc.Logger.Error("Error processing assets", "error", err)
		return nil, err
	}

	ids := getAssetIDs(output.DeletedAssets)
	event := &DeleteActEvent{
		EventID: uuid.New().String(),
	}

	// Produce the event to notify other services about the act deletion.
	if err := uc.Producer.Publish(ctx, event, e.DeleteActTopic); err != nil {
		uc.Logger.Error("Failed to produce event", "error", err)
		return nil, errors.NewInternal("error producing event", err)
	}

	// Return an empty output indicating successful deletion.
	return &DeleteActOutput{}, nil
}
