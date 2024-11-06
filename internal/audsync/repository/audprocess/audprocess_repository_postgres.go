package audprocess

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/audsync/entity"
	"github.com/JorgeO3/flowcast/pkg/postgres"
)

const (
	insertAudioProcessingQuery = ``
	getAudioProcessingQuery    = ``
	getAllAudioProcessingQuery = ``
	updateAudioProcessingQuery = ``
	deleteAudioProcessingQuery = ``
)

// PostgresAudprocessRepository implements the audprocess repository contract.
type PostgresAudprocessRepository struct {
	*postgres.Postgres
}

// NewRepository creates a new instance of PostgresAudprocessRepository.
func NewRepository(pg *postgres.Postgres) Repository {
	return &PostgresAudprocessRepository{pg}
}

// CreateProcess implements Repository.
func (p *PostgresAudprocessRepository) CreateProcess(ctx context.Context, audProcess *entity.AudioProcessing) error {
	panic("unimplemented")
}

// DeleteProcess implements Repository.
func (p *PostgresAudprocessRepository) DeleteProcess(ctx context.Context, eventID string) error {
	panic("unimplemented")
}

// GetAProcess implements Repository.
func (p *PostgresAudprocessRepository) GetAProcess(ctx context.Context, limit int, offset int) ([]*entity.AudioProcessing, error) {
	panic("unimplemented")
}

// GetProcess implements Repository.
func (p *PostgresAudprocessRepository) GetProcess(ctx context.Context, eventID string) (*entity.AudioProcessing, error) {
	panic("unimplemented")
}

// UpdateProcess implements Repository.
func (p *PostgresAudprocessRepository) UpdateProcess(ctx context.Context, audProcess *entity.AudioProcessing) error {
	panic("unimplemented")
}
