// Package assets implements the assets repository contract.
package assets

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/minio/minio-go/v7"
)

// RepositoryMinio implements the assets repository contract.
type RepositoryMinio struct {
	client *minio.Client
	bucket string
}

// NewRepository creates a new instance of RepositoryMinio.
func NewRepository(client *minio.Client) repository.AssetsRepository {
	return &RepositoryMinio{
		client: client,
	}
}

// GeneratePresignedURL implements RawAudioRepository.
func (r *RepositoryMinio) GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error) {
	url, err := r.client.PresignedPutObject(ctx, r.bucket, fileName, time)
	return url.String(), err
}
