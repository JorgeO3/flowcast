// Package repository package provides the different repositories for the catalog service
package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ActRepository represent the act repository contract
type ActRepository interface {
	CreateAct(context.Context, *entity.Act) error
	CreateManyActs(context.Context, []*entity.Act) error
	GetActByID(context.Context, primitive.ObjectID) (*entity.Act, error)
	GetManyActs(context.Context, bson.M, *options.FindOptions) ([]*entity.Act, error)
	UpdateAct(context.Context, *entity.Act) error
	UpdateManyActs(context.Context, bson.M, bson.M) error
	DeleteAct(context.Context, primitive.ObjectID) error
	DeleteManyActs(context.Context, bson.M) error
}