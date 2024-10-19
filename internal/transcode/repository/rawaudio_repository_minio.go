package repository

import (
	"context"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

// TempDir is the directory where the raw audio files are stored.
const TempDir string = "/tmp/minio_files"

// RawAudioRepositoryMinio is a repository for storing raw audio files in MinIO.
type RawAudioRepositoryMinio struct {
	client *minio.Client
	bucket string
}

// NewRawAudioRepository creates a new instance of RawAudioRepositoryMinio.
func NewRawAudioRepository(client *minio.Client, bucket string) RawAudioRepository {
	return &RawAudioRepositoryMinio{client: client, bucket: bucket}
}

// DownloadSong implements RawAudioRepository interface.
func (r *RawAudioRepositoryMinio) DownloadSong(ctx context.Context, songName string) (string, error) {
	opts := minio.GetObjectOptions{}
	filePath := filepath.Join(TempDir, songName)

	if err := os.MkdirAll(TempDir, os.ModePerm); err != nil {
		return "", err
	}

	if err := r.client.FGetObject(ctx, r.bucket, songName, filePath, opts); err != nil {
		return "", err
	}

	return filePath, nil
}
