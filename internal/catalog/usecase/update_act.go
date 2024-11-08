package usecase

import (
	"context"
	"time"

	e "github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
	"github.com/JorgeO3/flowcast/internal/catalog/infrastructure/kafka"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	"github.com/JorgeO3/flowcast/internal/catalog/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/mongotx"
	"github.com/JorgeO3/flowcast/pkg/redpanda"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"github.com/google/uuid"
)

// UpdateActInput represents the input required to update a musical act.
type UpdateActInput struct {
	Act    e.Act  `json:"act" validate:"required,dive"`
	UserID string `json:"userId" validate:"required"`
}

// UpdateActOutput represents the result of updating a musical act.
type UpdateActOutput struct {
	AssetURLs []utils.AssetURL `json:"assetUrls,omitempty"`
}

// BaseEvent defines the common structure for all events.
type BaseEvent struct {
	EventID   string    `json:"eventId"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

// UpdateActEvent represents an event triggered after updating an act.
type UpdateActEvent struct {
	BaseEvent
	DeletedAssets []utils.AssetChange `json:"deletedAssets"`
	AddedAssets   []utils.AssetChange `json:"addedAssets"`
}

// UpdateActUC is the use case responsible for updating a musical act.
type UpdateActUC struct {
	repos     *repository.Repositories
	logger    logger.Interface
	validator validator.Interface

	producer       *kafka.Producer
	transactionMgr mongotx.TxManager
}

// UpdateActUCOpt represents a functional option for configuring UpdateActUC.
type UpdateActUCOpt func(*UpdateActUC)

// optionHelper is a helper function to create functional options.
func optionHelper(setter func(*UpdateActUC)) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		setter(uc)
	}
}

// WithUpdateRepositories sets the repositories in the UpdateActUC.
func WithUpdateRepositories(repo *repository.Repositories) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.repos = repo
	}
}

// WithUpdateActLogger sets the logger in the UpdateActUC.
func WithUpdateActLogger(logger logger.Interface) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.logger = logger
	}
}

// WithUpdateActValidator sets the Validator in UpdateActUC.
func WithUpdateActValidator(v validator.Interface) UpdateActUCOpt {
	return optionHelper(func(uc *UpdateActUC) {
		uc.validator = v
	})
}

// WithUpdateActProducer sets the Producer in UpdateActUC.
func WithUpdateActProducer(producer redpanda.Producer) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.producer = producer
	}
}

// WithUpdateActTransactionManager sets the TransactionManager in UpdateActUC.
func WithUpdateActTransactionManager(txMgr mongotx.TxManager) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.transactionMgr = txMgr
	}
}

// NewUpdateAct creates a new instance of UpdateActUC with the provided options.
func NewUpdateAct(opts ...UpdateActUCOpt) *UpdateActUC {
	uc := &UpdateActUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

func (uc *UpdateActUC) handleUpdateEvent(userID string, addedAssets, deletedAssets []utils.AssetChange) error {
	event := UpdateActEvent{
		BaseEvent: BaseEvent{
			EventID:   uuid.New().String(),
			UserID:    userID,
			CreatedAt: time.Now(),
		},
		DeletedAssets: deletedAssets,
		AddedAssets:   addedAssets,
	}

	if err := uc.producer.Publish(event, e.UpdateActTopic); err != nil {
		uc.logger.Error("Failed to publish event: %v", err)
		return errors.NewInternal("failed to publish event", err)
	}

	return nil
}

func (uc *UpdateActUC) updateActTransaction(ctx context.Context, newAct *e.Act, ucOutput *UpdateActOutput) error {
	oldAct, err := uc.repos.Act.GetActByID(ctx, newAct.ID)
	if err != nil {
		return errors.HandleRepoError(err)
	}

	// * The Processor works as follows:
	// * 1. Identify the changes in the assets comparing the old and new acts.
	// * 2. Process all the changes categorizing them as added, deleted or updated assets.
	// *  - Exists 4 types of assets: Audio, ImageAct, ImageAlbum, ImageSong.
	// *  - For each type of asset the process is the following:
	// * 	- For the added assets, generate the URLs.
	// * 	- For the deleted assets, remove the assets from the storage.
	// * 	- For the updated assets, remove the old assets and generate the URLs for the new assets.
	// * 3. Return the output with all the necesary information for the client and audio service.
	params := utils.NewAssetsProcessorParams(ctx, oldAct, newAct, uc.repos)
	processor := utils.NewAssetsProcessor(params)
	processor.IdentifyChanges()
	output, err := processor.ProcessChanges()
	if err != nil {
		uc.logger.Error("Failed to process asset changes: %v", err)
		return err
	}

	if err := uc.repos.Act.UpdateAct(ctx, newAct); err != nil {
		return errors.HandleRepoError(err)
	}

	if err := uc.handleUpdateEvent(newAct.UserID, output.GetAddedAssets(), output.GetDeletedAssets()); err != nil {
		return err
	}

	ucOutput.AssetURLs = output.GetAssetsURLs()
	return nil
}

// Execute performs the UpdateAct use case, updating the act and handling associated changes.
func (uc *UpdateActUC) Execute(ctx context.Context, input UpdateActInput) (*UpdateActOutput, error) {
	uc.logger.Info("Starting act update in the catalog")

	// Validate the input
	if err := uc.validator.Validate(input); err != nil {
		uc.logger.Warn("Invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Execute the transaction to update the act
	var output UpdateActOutput
	err := uc.transactionMgr.Run(ctx, func(ctx context.Context) error {
		return uc.updateActTransaction(ctx, &input.Act, &output)
	})
	if err != nil {
		uc.logger.Error("Act update failed: %v", err)
		return nil, err
	}

	uc.logger.Info("Act updated successfully")
	return &output, nil
}
