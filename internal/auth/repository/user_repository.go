// Package repository provides the repository for the auth service.
package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/auth/entity"
)

// UserRepository is the interface that defines the methods for the user repository.
type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByID(ctx context.Context, id int) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) (int, error)
	Update(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	LockUser(ctx context.Context, id int) error
}
