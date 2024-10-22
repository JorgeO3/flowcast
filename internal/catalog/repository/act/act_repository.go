// Package act repository  package provides the different repositories for the catalog service
package act

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
)

// Repository represent the act repository contract
type Repository interface {
	CreateAct(context.Context, *entity.Act) (string, error)
	CreateActs(context.Context, []*entity.Act) ([]string, error)
	UpdateAct(context.Context, *entity.Act) error
	DeleteAct(context.Context, string) error
	GetActByID(context.Context, string) (*entity.Act, error)
	GetActs(context.Context, string, int64, int64) ([]*entity.Act, error)
}
