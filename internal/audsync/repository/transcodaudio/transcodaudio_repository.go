package transcodaudio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type MinioRepository struct {
	client *minio.Client
	bucket string
}

// DeleteObject implements Repository.
func (m *MinioRepository) DeleteObject(ctx context.Context, filePath string) error {
	panic("unimplemented")
}

// DeleteObjects implements Repository.
func (m *MinioRepository) DeleteObjects(ctx context.Context, filePaths []string) error {
	panic("unimplemented")
}

// NewRepository creates a new instance of MinioRepository.
func NewRepository(client *minio.Client, bucket string) Repository {
	return &MinioRepository{
		client: client,
		bucket: bucket,
	}
}
