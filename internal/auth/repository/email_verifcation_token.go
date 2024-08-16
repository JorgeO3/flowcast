package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
)

// EmailVerificationTokenRepo -.
type EmailVerificationTokenRepo interface {
	FindByUserID(ctx context.Context, userID int) (*entity.EmailVerificationToken, error)
}
