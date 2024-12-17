package usecase

import (
	"context"

	e "github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/errors"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/eventbus"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/logger"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/validator"
	"github.com/JorgeO3/flowcast/pkg/mongotx"
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
	Logger       logger.Interface
	Validator    validator.Interface
	Producer     eventbus.Producer
	ActRepo      repository.ActRepository
	AssetsRepo   repository.AssetsRepository
	RawaudioRepo repository.RawaudioRepository
	tx           *mongotx.MongoTx
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
func WithDeleteActProducer(prod eventbus.Producer) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.Producer = prod
	}
}

// WithDeleteActActRepo injects the ActRepository into the use case.
// It allows the use case to access and manipulate act data in the repository.
func WithDeleteActActRepo(repo repository.ActRepository) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.ActRepo = repo
	}
}

// WithDeleteActAssetsRepo injects the AssetsRepository into the use case.
// It allows the use case to access and manipulate asset data in the repository.
func WithDeleteActAssetsRepo(repo repository.AssetsRepository) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.AssetsRepo = repo
	}
}

// WithDeleteActRawaudioRepo injects the RawaudioRepository into the use case.
// It allows the use case to access and manipulate rawaudio data in the repository.
func WithDeleteActRawaudioRepo(repo repository.RawaudioRepository) DeleteActOpts {
	return func(uc *DeleteActUC) {
		uc.RawaudioRepo = repo
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
	act, err := uc.ActRepo.ReadOne(ctx, id)
	if err != nil {
		uc.Logger.Errorf("failed to read act from repository: %v", err)
		return nil, err
	}

	// Attempt to delete the act from the repository using the provided ID.
	if err := uc.ActRepo.DeleteOne(ctx, id); err != nil {
		uc.Logger.Errorf("failed to delete act from repository: %v", err)
		return nil, err
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
	uc.Logger.Infof("deleting musical act from the catalog")

	// Validate input parameters to ensure required fields are present and correct.
	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Warnf("invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	var act *e.Act
	var err error
	err = uc.tx.Run(ctx, func(ctx context.Context) error {
		act, err = uc.handleDelete(ctx, input.ID)
		return err
	})

	if err != nil {
		uc.Logger.Errorf("error deleting act: %v", err)
		return nil, err
	}

	processor := NewAssetsProcessor(ctx, uc.ActRepo, uc.RawaudioRepo, uc.AssetsRepo)
	_, err = processor.Delete(act)
	if err != nil {
		uc.Logger.Errorf("error processing assets: %v", err)
		return nil, err
	}

	// FIXME: Implement the publisher correctly
	event := struct{}{}

	// Produce the event to notify other services about the act deletion.
	if err := uc.Producer.Publish(ctx, event, e.DeleteActTopic); err != nil {
		uc.Logger.Errorf("error producing event: %v", err)
		return nil, err
	}

	// Return an empty output indicating successful deletion.
	return &DeleteActOutput{}, nil
}
