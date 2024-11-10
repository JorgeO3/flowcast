// Package repository provides the different repositories for the catalog service
package repository

import (
	"context"
	"time"
)

// AssetsRepository respresents the assets repository contract
type AssetsRepository interface {
	GeneratePresignedURL(ctx context.Context, fileName string, time time.Duration) (string, error)
}
