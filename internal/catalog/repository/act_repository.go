// Package repository package provides the different repositories for the catalog service
package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/catalog/entity"
)

// ActRepository represent the act repository contract
type ActRepository interface {
	CreateAct(ctx context.Context, act entity.Act) error
	GetActByID(ctx context.Context, id string) (*entity.Act, error)
	UpdateAct(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteAct(ctx context.Context, id string) error

	// Operaciones de alto nivel
	AddMember(ctx context.Context, actID string, member entity.Member) error
	RemoveMember(ctx context.Context, actID, memberID string) error
	AddAlbum(ctx context.Context, actID string, album entity.Album) error
	RemoveAlbum(ctx context.Context, actID, albumID string) error
}
