package repository

import (
	"context"
	"fmt"
	"strings"

	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/pkg/postgres"
	"gitlab.com/JorgeO3/flowcast/pkg/transaction"
)

// PostgresSocialLinkRepo implements SocialLinkRepo for PostgreSQL.
type PostgresSocialLinkRepo struct {
	*postgres.Postgres
}

// NewPostgresSocialLinkRepo creates a new instance of PostgresSocialLinkRepo.
func NewPostgresSocialLinkRepo(db *postgres.Postgres) SocialLinkRepo {
	return &PostgresSocialLinkRepo{db}
}

const (
	getSocialLinksQuery = `
	SELECT id, user_id, platform, url 
	FROM social_links
	WHERE user_id = $1;
	`

	updateSocialLinksQuery = `
	UPDATE social_links
	SET url = $1, platform = $2
	WHERE id = $3;
	`

	insertSocialLinksHeaderQuery = `
	INSERT INTO social_links 
	(
		user_id, 
		platform, 
		url
	) 
	VALUES %s;
	`
)

// FindByUserID retrieves a social link by user ID.
func (p *PostgresSocialLinkRepo) FindByUserID(ctx context.Context, userID int) (*entity.SocialLink, error) {
	var socialLink entity.SocialLink
	dest := []interface{}{
		&socialLink.ID,
		&socialLink.UserID,
		&socialLink.Platform,
		&socialLink.URL,
	}

	err := p.Pool.QueryRow(ctx, getSocialLinksQuery, userID).Scan(dest...)
	return &socialLink, err
}

// FindByUserIDTx retrieves a social link by user ID within a transaction.
func (p *PostgresSocialLinkRepo) FindByUserIDTx(ctx context.Context, tx transaction.Tx, userID int) (*entity.SocialLink, error) {
	var socialLink entity.SocialLink
	dest := []interface{}{
		&socialLink.ID,
		&socialLink.UserID,
		&socialLink.Platform,
		&socialLink.URL,
	}

	err := tx.QueryRow(ctx, getSocialLinksQuery, userID).Scan(dest...)
	return &socialLink, err
}

// SaveTx saves multiple social links within a transaction.
func (p *PostgresSocialLinkRepo) SaveTx(ctx context.Context, tx transaction.Tx, socialLinks []*entity.SocialLink) error {
	if len(socialLinks) == 0 {
		return nil // No social links to save
	}

	args, valueStrings := buildSocialLinkInsertArgs(socialLinks)
	insertSocialLinksQuery := fmt.Sprintf(insertSocialLinksHeaderQuery, strings.Join(valueStrings, ","))

	_, err := tx.Exec(ctx, insertSocialLinksQuery, args...)
	return err
}

// Save saves multiple social links.
func (p *PostgresSocialLinkRepo) Save(ctx context.Context, socialLinks []*entity.SocialLink) error {
	if len(socialLinks) == 0 {
		return nil // No social links to save
	}

	args, valueStrings := buildSocialLinkInsertArgs(socialLinks)
	insertSocialLinksQuery := fmt.Sprintf(insertSocialLinksHeaderQuery, strings.Join(valueStrings, ","))

	_, err := p.Pool.Exec(ctx, insertSocialLinksQuery, args...)
	return err
}

// Update updates a social link.
func (p *PostgresSocialLinkRepo) Update(ctx context.Context, socialLink *entity.SocialLink) error {
	args := []interface{}{
		socialLink.URL,
		socialLink.Platform,
		socialLink.ID,
	}

	_, err := p.Pool.Exec(ctx, updateSocialLinksQuery, args...)
	return err
}

// buildSocialLinkInsertArgs constructs the arguments and value strings for inserting social links.
func buildSocialLinkInsertArgs(socialLinks []*entity.SocialLink) ([]interface{}, []string) {
	args := []interface{}{}
	valueStrings := []string{}
	for i, socialLink := range socialLinks {
		n := i * 3
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", n+1, n+2, n+3))
		args = append(args, socialLink.UserID, socialLink.Platform, socialLink.URL)
	}
	return args, valueStrings
}
