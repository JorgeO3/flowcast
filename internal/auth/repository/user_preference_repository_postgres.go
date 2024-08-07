package repository

import (
	"context"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
	"gitlab.com/JorgeO3/flowcast/pkg/txmanager"
)

// PostgresUserPrefRepo is the implementation of the user preference repository for PostgreSQL.
type PostgresUserPrefRepo struct {
	*postgres.Postgres
}

// NewPostgresUserPrefRepo creates a new instance of PostgresUserPrefRepo.
func NewPostgresUserPrefRepo(db *postgres.Postgres) UserPreferenceRepository {
	return &PostgresUserPrefRepo{db}
}

const (
	searchUserPrefQuery = `
	SELECT id, user_id, email_notifications, sms_notifications
	FROM user_preferences
	WHERE user_id = $1;
	`

	insertUserPrefQuery = `
	INSERT INTO user_preferences
	(
		user_id,
		email_notifications,
		sms_notifications
	)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	updateUserPrefQuery = `
	UPDATE user_preferences
	SET
		email_notifications = $1,
		sms_notifications = $2
	WHERE user_id = $3;
	`
)

// FindByUserID searches for a user preference by their user ID.
func (p *PostgresUserPrefRepo) FindByUserID(ctx context.Context, userID int) (*entity.UserPref, error) {
	var userPref entity.UserPref

	dest := []interface{}{
		&userPref.ID,
		&userPref.UserID,
		&userPref.EmailNotifications,
		&userPref.SmsNotifications,
	}

	db := txmanager.GetTx(ctx, p.Pool)
	err := db.QueryRow(ctx, searchUserPrefQuery, userID).Scan(dest...)
	return &userPref, err
}

// Save saves a user preference.
func (p *PostgresUserPrefRepo) Save(ctx context.Context, userPref *entity.UserPref) error {
	args := []interface{}{
		userPref.UserID,
		userPref.EmailNotifications,
		userPref.SmsNotifications,
	}

	db := txmanager.GetTx(ctx, p.Pool)
	_, err := db.Exec(ctx, insertUserPrefQuery, args...)
	return err
}

// Update updates a user preference.
func (p *PostgresUserPrefRepo) Update(ctx context.Context, userPref *entity.UserPref) error {
	args := []interface{}{
		userPref.EmailNotifications,
		userPref.SmsNotifications,
		userPref.UserID,
	}

	db := txmanager.GetTx(ctx, p.Pool)
	_, err := db.Exec(ctx, updateUserPrefQuery, args...)
	return err
}
