// Package audprocess provides the repository for the audprocess service
package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/entity"
)

// Repository is an interface for the audprocess repository.
type Repository interface {
	CreateProcess(ctx context.Context, audProcess *entity.AudioProcessing) error
	UpdateProcess(ctx context.Context, audProcess *entity.AudioProcessing) error
	GetProcess(ctx context.Context, eventID string) (*entity.AudioProcessing, error)
	GetAProcess(ctx context.Context, limit, offset int) ([]*entity.AudioProcessing, error)
	DeleteProcess(ctx context.Context, eventID string) error
}
