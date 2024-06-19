package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// UserRepo is the interface that defines the methods for the user repository.
type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindByID(ctx context.Context, id int) (*entity.User, error)
	Save(ctx context.Context, tx transaction.Tx, user *entity.User) (int, error)
	Update(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	LockUser(ctx context.Context, id int) error
}
