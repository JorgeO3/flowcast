package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

const (
	insertEmailVerificationTokenQuery = `
	INSERT INTO email_verification_tokens 
	(
		user_id,
		token,
		created_at
	)
	VALUES ($1, $2, $3)
	RETURNING id;
	`
	getEmailVerificationTokenQuery = `
	SELECT id, user_id, token, created_at
	FROM email_verification_tokens
	WHERE user_id = $1;
	`
)

// PostgresEmailVerificationTokenRepo -.
type PostgresEmailVerificationTokenRepo struct {
	*postgres.Postgres
}

// FindByUserID implements EmailVerificationTokenRepo.
func (p PostgresEmailVerificationTokenRepo) FindByUserID(ctx context.Context, userID int) (*entity.EmailVerificationToken, error) {
	var emailVerificationT entity.EmailVerificationToken

	dest := []interface{}{
		&emailVerificationT.ID,
		&emailVerificationT.UserID,
		&emailVerificationT.Token,
		&emailVerificationT.CreatedAt,
	}

	err := p.Pool.QueryRow(ctx, getEmailVerificationTokenQuery, userID).Scan(dest...)
	return &emailVerificationT, postgres.MapError(err)
}

// SaveTx implements EmailVerificationTokenRepo.
func (p PostgresEmailVerificationTokenRepo) SaveTx(ctx context.Context, tx transaction.Tx, token *entity.EmailVerificationToken) error {
	args := []interface{}{
		token.UserID,
		token.Token,
		token.CreatedAt,
	}

	_, err := tx.Exec(ctx, insertEmailVerificationTokenQuery, args...)
	return postgres.MapError(err)
}

// NewPostgresEmailVerificationTokenRepo -.
func NewPostgresEmailVerificationTokenRepo(pg *postgres.Postgres) EmailVerificationTokenRepo {
	return &PostgresEmailVerificationTokenRepo{pg}
}
