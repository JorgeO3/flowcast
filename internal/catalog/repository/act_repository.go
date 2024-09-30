// Package repository package provides the different repositories for the catalog service
package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActRepository represent the act repository contract
type ActRepository interface {
	CreateAct(context.Context, *entity.Act) (string, error)
	CreateManyActs(context.Context, []*entity.Act) ([]string, error)
	UpdateAct(context.Context, *entity.Act) error
	DeleteAct(context.Context, primitive.ObjectID) error
	GetActByID(context.Context, primitive.ObjectID) (*entity.Act, error)
	GetActs(context.Context, string, int64, int64) ([]*entity.Act, error)
}
