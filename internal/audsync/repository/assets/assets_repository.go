// Package assets provides all the methods to interact with the assets bucket.
package assets

import "context"

// Repository is an interface for the assets repository.
type Repository interface {
	DeleteObject(ctx context.Context, filePath string) error
	DeleteObjects(ctx context.Context, filePaths []string) error
}
