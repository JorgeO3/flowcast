package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// EmailVerificationTokenRepo -.
type EmailVerificationTokenRepo interface {
	SaveTx(ctx context.Context, tx transaction.Tx, token *entity.EmailVerificationToken) error
	FindByUserID(ctx context.Context, userID int) (*entity.EmailVerificationToken, error)
}
