package usecase

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/errors"
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/eventbus"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/logger"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/utils"
	"github.com/JorgeO3/flowcast/internal/catalogv2/usecase/validator"
	"github.com/JorgeO3/flowcast/pkg/mongotx"
	"github.com/google/uuid"
)

// UpdateActInput represents the input required to update a musical act.
type UpdateActInput struct {
	Act    entity.Act `json:"act" validate:"required,dive"`
	UserID string     `json:"userId" validate:"required"`
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
	DeletedAssets []AssetChange `json:"deletedAssets"`
	AddedAssets   []AssetChange `json:"addedAssets"`
}

// UpdateActUC is the use case responsible for updating a musical act.
type UpdateActUC struct {
	logger    logger.Interface
	producer  eventbus.Producer
	validator validator.Interface

	transactionMgr mongotx.TxManager
	actRepo        repository.ActRepository
	assetsRepo     repository.AssetsRepository
	rawaudioRepo   repository.RawaudioRepository
}

// UpdateActUCOpt represents a functional option for configuring UpdateActUC.
type UpdateActUCOpt func(*UpdateActUC)

// WithUpdateActRepository sets the repository in the UpdateActUC.
func WithUpdateActRepository(repo repository.ActRepository) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.actRepo = repo
	}
}

// WithUpdateActAssetsRepository sets the repository in the UpdateActUC.
func WithUpdateActAssetsRepository(repo repository.AssetsRepository) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.assetsRepo = repo
	}
}

// WithUpdateActRawaudioRepository sets the repository in the UpdateActUC.
func WithUpdateActRawaudioRepository(repo repository.RawaudioRepository) UpdateActUCOpt {
	return func(uc *UpdateActUC) {
		uc.rawaudioRepo = repo
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
	return func(uc *UpdateActUC) {
		uc.validator = v
	}
}

// WithUpdateActProducer sets the Producer in UpdateActUC.
func WithUpdateActProducer(producer eventbus.Producer) UpdateActUCOpt {
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

func (uc *UpdateActUC) handleUpdateEvent(ctx context.Context, userID string, addedAssets, deletedAssets []AssetChange) error {
	event := UpdateActEvent{
		BaseEvent: BaseEvent{
			EventID:   uuid.New().String(),
			UserID:    userID,
			CreatedAt: time.Now(),
		},
		DeletedAssets: deletedAssets,
		AddedAssets:   addedAssets,
	}

	if err := uc.producer.Publish(ctx, event, entity.UpdateActTopic); err != nil {
		uc.logger.Errorf("failed to publish event: %v", err)
		return err
	}

	return nil
}

func (uc *UpdateActUC) updateActTransaction(ctx context.Context, newAct *entity.Act, ucOutput *UpdateActOutput) error {
	oldAct, err := uc.actRepo.ReadOne(ctx, newAct.ID)
	if err != nil {
		uc.logger.Errorf("failed to read act from repository: %v", err)
		return err
	}

	if err := uc.actRepo.UpdateOne(ctx, newAct); err != nil {
		uc.logger.Errorf("error updating act: %v", err)
		return err
	}

	processor := NewAssetsProcessor(ctx, uc.actRepo, uc.rawaudioRepo, uc.assetsRepo)
	output, err := processor.Update(oldAct, newAct)
	if err != nil {
		uc.logger.Errorf("error processing assets: %v", err)
		return err
	}

	// TODO: Handle the event

	return nil
}

// Execute performs the UpdateAct use case, updating the act and handling associated changes.
func (uc *UpdateActUC) Execute(ctx context.Context, input UpdateActInput) (*UpdateActOutput, error) {
	uc.logger.Info("Starting act update in the catalog")

	// Validate the input
	if err := uc.validator.Validate(input); err != nil {
		uc.logger.Warnf("invalid input: %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	// Execute the transaction to update the act
	var output UpdateActOutput
	err := uc.transactionMgr.Run(ctx, func(ctx context.Context) error {
		return uc.updateActTransaction(ctx, &input.Act, &output)
	})
	if err != nil {
		uc.logger.Errorf("error updating act: %v", err)
		return nil, err
	}

	uc.logger.Info("Act updated successfully")
	return &output, nil
}
