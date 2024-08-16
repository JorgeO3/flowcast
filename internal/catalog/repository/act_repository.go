// Package repository package provides the different repositories for the catalog service
package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
)

// ActRepository represents the repository for acts.
type ActRepository interface {
	CreateAct(ctx context.Context, act *entity.Act) (*entity.Act, error)
	GetActByID(id string) (*entity.Act, error)
	UpdateAct(id string, updates map[string]interface{}) error
	DeleteAct(id string) error
}
