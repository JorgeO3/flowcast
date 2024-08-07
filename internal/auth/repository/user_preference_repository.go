package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
)

// UserPreferenceRepository is the interface for the user preference repository.
type UserPreferenceRepository interface {
	FindByUserID(ctx context.Context, userID int) (*entity.UserPref, error)
	Save(ctx context.Context, userPref *entity.UserPref) error
	Update(ctx context.Context, userPref *entity.UserPref) error
}
