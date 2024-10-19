// Package assets provides a repository for static assets.
package assets

import (
	"context"
	"time"
)

// Repository respresents the assets repository contract
type Repository interface {
	GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error)
}
