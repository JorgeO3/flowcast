package repository

import (
	"context"

	"github.com/JorgeO3/flowcast/internal/auth/entity"
	"github.com/JorgeO3/flowcast/pkg/postgres"
	"github.com/JorgeO3/flowcast/pkg/txmanager"
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

// NewPostgresEmailVerificationTokenRepo -.
func NewPostgresEmailVerificationTokenRepo(pg *postgres.Postgres) PostgresEmailVerificationTokenRepo {
	return PostgresEmailVerificationTokenRepo{pg}
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

	db := txmanager.GetTx(ctx, p.Pool)
	err := db.QueryRow(ctx, getEmailVerificationTokenQuery, userID).Scan(dest...)
	return &emailVerificationT, err
}
