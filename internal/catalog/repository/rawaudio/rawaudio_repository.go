package rawaudio

import (
	"context"
	"time"
)

// Repository respresents the raw audio repository contract
type Repository interface {
	GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error)
}
