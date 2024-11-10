package repository

import (
	"context"
	"time"
)

// RawaudioRepository respresents the raw audio repository contract
type RawaudioRepository interface {
	GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error)
}
