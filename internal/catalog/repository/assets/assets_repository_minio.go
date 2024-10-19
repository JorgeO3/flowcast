package assets

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

// RepositoryMinio implements the assets repository contract.
type RepositoryMinio struct {
	client *minio.Client
	bucket string
}

// NewRepository creates a new instance of RepositoryMinio.
func NewRepository(client *minio.Client) Repository {
	return &RepositoryMinio{
		client: client,
	}
}

// GeneratePresignedURL implements RawAudioRepository.
func (r *RepositoryMinio) GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error) {
	url, err := r.client.PresignedPutObject(ctx, r.bucket, fileName, time)
	return url.String(), err
}
