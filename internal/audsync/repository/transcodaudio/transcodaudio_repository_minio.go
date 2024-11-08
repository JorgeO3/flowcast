package transcodaudio

import "context"

type Repository interface {
	DeleteObject(ctx context.Context, filePath string) error
	DeleteObjects(ctx context.Context, filePaths []string) error
}
