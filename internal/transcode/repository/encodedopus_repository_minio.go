package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

// EncodedOpusRepositoryMinio is a repository for storing encoded Opus files in MinIO.
type EncodedOpusRepositoryMinio struct {
	client *minio.Client
	bucket string
}

// NewEncodedOpusRepository creates a new instance of EncodedOpusRepositoryMinio.
func NewEncodedOpusRepository(client *minio.Client, bucket string) EncodedOpusRepository {
	return &EncodedOpusRepositoryMinio{client: client, bucket: bucket}
}

// PutSong implements EncodedOpusRepository.
func (e *EncodedOpusRepositoryMinio) PutSong(ctx context.Context, songName string, outputDir string) error {
	filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		relativePath, _ := filepath.Rel(outputDir, path)
		minioPath := filepath.Join(songName, relativePath)
		opts := minio.PutObjectOptions{}

		_, err = e.client.FPutObject(ctx, e.bucket, minioPath, path, opts)
		return err
	})

	return nil
}
