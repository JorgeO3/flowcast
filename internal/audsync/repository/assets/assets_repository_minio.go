package assets

import (
	"context"

	"github.com/minio/minio-go/v7"
)

// MinioRepository implements the assets repository contract.
type MinioRepository struct {
	client *minio.Client
	bucket string
}

// NewRepository creates a new instance of MinioRepository.
func NewRepository(client *minio.Client, bucket string) Repository {
	return &MinioRepository{
		client: client,
		bucket: bucket,
	}
}

// DeleteObject implements Repository.
func (m *MinioRepository) DeleteObject(ctx context.Context, filePath string) error {
	opts := minio.RemoveObjectOptions{}
	return m.client.RemoveObject(ctx, m.bucket, filePath, opts)
}

// DeleteObjects implements Repository.
func (m *MinioRepository) DeleteObjects(ctx context.Context, filePaths []string) error {
	panic("unimplemented")
}
