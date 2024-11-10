// Package rawaudio provides the different repositories for the catalog service
package rawaudio

import (
	"context"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/repository"
	"github.com/minio/minio-go/v7"
)

// RepositoryMinio is a repository for the raw audio files.
type RepositoryMinio struct {
	client *minio.Client
	bucket string
}

// NewRepository creates a new instance of RawAudioRepositoryMinio.
func NewRepository(client *minio.Client, bucket string) repository.RawaudioRepository {
	return &RepositoryMinio{
		client: client,
		bucket: bucket,
	}
}

// GeneratePresignedURL implements RawAudioRepository.
func (r *RepositoryMinio) GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error) {
	url, err := r.client.PresignedPutObject(ctx, r.bucket, fileName, time)
	return url.String(), err
}
