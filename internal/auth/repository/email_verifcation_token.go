package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
)

// EmailVerificationTokenRepo -.
type EmailVerificationTokenRepo interface {
	SaveTx(ctx context.Context, token *entity.EmailVerificationToken) error
	FindByUserID(ctx context.Context, userID int) (*entity.EmailVerificationToken, error)
}
