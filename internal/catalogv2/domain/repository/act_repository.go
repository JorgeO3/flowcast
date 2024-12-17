// Package repository provides the different repositories for the catalog service
package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
)

// ActRepository represent the act repository contract
type ActRepository interface {
	CreateOne(context.Context, *entity.Act) (string, error)
	CreateMany(context.Context, []entity.Act) ([]string, error)
	UpdateOne(context.Context, *entity.Act) error
	DeleteOne(context.Context, string) error
	ReadOne(context.Context, string) (*entity.Act, error)
	ReadMany(context.Context, string, int64, int64) ([]*entity.Act, error)
}
