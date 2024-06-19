// Package repository contains the implementation of the user repository for PostgreSQL.
package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
)

// PostgresUserRepo is the implementation of the user repository for PostgreSQL.
type PostgresUserRepo struct {
	*postgres.Postgres
}

// NewPostgresUserRepo creates a new instance of PostgresUserRepo.
func NewPostgresUserRepo(db *postgres.Postgres) *PostgresUserRepo {
	return &PostgresUserRepo{db}
}

// FindByUsername searches for a user by their username.
func (p *PostgresUserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	tx, err := p.Pool.Begin(ctx)
	return &entity.User{}, nil
}

// FindByID searches for a user by their ID.
func (p *PostgresUserRepo) FindByID(ctx context.Context, id int) (*entity.User, error) {
	return &entity.User{}, nil
}

const insertUserQuery = `
	INSERT INTO users
	(
		username,
		email,
		password,
		full_name,
		birth_date,
		gender,
		phone_number,
		status,
		subscription_status,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
	ON CONFLICT DO NOTHING
`

// Save -.
func (p *PostgresUserRepo) Save(ctx context.Context, user *entity.User) error {
	args := []any{
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.Birthdate,
		user.Gender,
		user.Phone,
		user.AuthStatus,
		user.SubscriptionStatus,
		user.CreatedAt,
		user.UpdatedAt,
	}
	_, err := p.Pool.Exec(ctx, insertUserQuery, args...)
	return err
}

// Update updates an existing user in the database.
func (p *PostgresUserRepo) Update(ctx context.Context, user *entity.User) error {
	return nil
}

// FindByEmail searches for a user by their email.
func (p *PostgresUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return &entity.User{}, nil
}

// LockUser locks a user by their ID.
func (p *PostgresUserRepo) LockUser(ctx context.Context, id int) error {
	return nil
}
