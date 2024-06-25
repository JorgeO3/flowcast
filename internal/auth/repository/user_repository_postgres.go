// Package repository contains the implementation of the user repository for PostgreSQL.
package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// PostgresUserRepo is the implementation of the user repository for PostgreSQL.
type PostgresUserRepo struct {
	*postgres.Postgres
}

// NewPostgresUserRepo creates a new instance of PostgresUserRepo.
func NewPostgresUserRepo(db *postgres.Postgres) UserRepo {
	return &PostgresUserRepo{db}
}

const (
	searchUserQuery = `
	SELECT
		id, username, email, password, full_name, birth_date, gender,
		phone_number, status, subscription_status, created_at, updated_at
	FROM
		users
	WHERE
		username = $1;
	`

	searchUserByIDQuery = `
	SELECT
		id, username, email, password, full_name, birth_date, gender,
		phone_number, status, subscription_status, created_at, updated_at
	FROM
		users
	WHERE
		id = $1;
	`

	insertUserQuery = `
	INSERT INTO users
	(
		username, email, password, full_name, birth_date, gender, phone_number,
		status, subscription_status, created_at, updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id;
	`

	updateUserQuery = `
	UPDATE users
	SET
		username = $1, email = $2, password = $3, full_name = $4, birth_date = $5,
		gender = $6, phone_number = $7, status = $8, subscription_status = $9,
		created_at = $10, updated_at = $11
	WHERE id = $12;
	`

	searchUserByEmailQuery = `
	SELECT
		id, username, email, password, full_name, birth_date, gender,
		phone_number, status, subscription_status, created_at, updated_at
	FROM
		users
	WHERE
		email = $1;
	`

	lockUserQuery = `
	UPDATE users
	SET status = 'locked'
	WHERE id = $1;
	`
)

// FindByUsername searches for a user by their username.
func (p *PostgresUserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	return p.findUser(ctx, searchUserQuery, username)
}

// FindByID searches for a user by their ID.
func (p *PostgresUserRepo) FindByID(ctx context.Context, userID int) (*entity.User, error) {
	return p.findUser(ctx, searchUserByIDQuery, userID)
}

// FindByEmail searches for a user by their email.
func (p *PostgresUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return p.findUser(ctx, searchUserByEmailQuery, email)
}

func (p *PostgresUserRepo) findUser(ctx context.Context, query string, arg interface{}) (*entity.User, error) {
	var user entity.User

	dest := []interface{}{
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Birthdate,
		&user.Gender,
		&user.Phone,
		&user.Status,
		&user.SubsStatus,
		&user.CreatedAt,
		&user.UpdatedAt,
	}

	err := p.Pool.QueryRow(ctx, query, arg).Scan(dest...)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Return nil if no user is found
		}
		return nil, err // Return the error if something went wrong
	}

	return &user, nil
}

// Save saves a user in the database.
func (p *PostgresUserRepo) Save(ctx context.Context, tx transaction.Tx, user *entity.User) (int, error) {
	var userID int

	args := []interface{}{
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.Birthdate,
		user.Gender,
		user.Phone,
		user.Status,
		user.SubsStatus,
		user.CreatedAt,
		user.UpdatedAt,
	}
	err := tx.QueryRow(ctx, insertUserQuery, args...).Scan(&userID)
	return userID, postgres.MapError(err)
}

// Update updates an existing user in the database.
func (p *PostgresUserRepo) Update(ctx context.Context, user *entity.User) error {
	args := []interface{}{
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.Birthdate,
		user.Gender,
		user.Phone,
		user.Status,
		user.SubsStatus,
		user.CreatedAt,
		user.UpdatedAt,
		user.ID,
	}
	_, err := p.Pool.Exec(ctx, updateUserQuery, args...)
	return postgres.MapError(err)
}

// LockUser locks a user by their ID.
func (p *PostgresUserRepo) LockUser(ctx context.Context, id int) error {
	_, err := p.Pool.Exec(ctx, lockUserQuery, id)
	return postgres.MapError(err)
}
