package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// UserPrefRepo is the interface that defines the methods for the user preference repository.
type UserPrefRepo interface {
	FindByUserID(ctx context.Context, userID int) (*entity.UserPref, error)
	Save(ctx context.Context, tx transaction.Tx, userPref *entity.UserPref) error
	Update(ctx context.Context, userPref *entity.UserPref) error
}
