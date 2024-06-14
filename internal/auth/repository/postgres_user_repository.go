package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
)

// PostgresUserRepo is the implementation of the user repository for PostgreSQL.
type PostgresUserRepo struct {
	db *postgres.Postgres
}

// NewPostgresUserRepo creates a new instance of PostgresUserRepo.
func NewPostgresUserRepo(db *postgres.Postgres) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

// FindByUsername searches for a user by their username.
func (p *PostgresUserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	return &entity.User{}, nil
}

// FindByID searches for a user by their ID.
func (p *PostgresUserRepo) FindByID(ctx context.Context, id int) (*entity.User, error) {
	return &entity.User{}, nil
}

// Save saves a new user to the database.
func (p *PostgresUserRepo) Save(ctx context.Context, user *entity.User) error {
	return nil
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
