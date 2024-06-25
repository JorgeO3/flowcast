package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// SocialLinkRepo -.
type SocialLinkRepo interface {
	FindByUserID(ctx context.Context, userID int) (*entity.SocialLink, error)
	FindByUserIDTx(ctx context.Context, tx transaction.Tx, userID int) (*entity.SocialLink, error)
	SaveTx(ctx context.Context, tx transaction.Tx, socialLinks []*entity.SocialLink) error
	Save(ctx context.Context, socialLinks []*entity.SocialLink) error
	Update(ctx context.Context, socialLink *entity.SocialLink) error
}
