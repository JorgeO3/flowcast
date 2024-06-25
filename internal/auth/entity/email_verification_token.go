package entity

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

// EmailVerificationToken represents an email verification token entity.
type EmailVerificationToken struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
}

// NewEmailVerificationToken creates a new instance of EmailVerficationToken.
func NewEmailVerificationToken(userID int) (*EmailVerificationToken, error) {
	token, err := generateToken()
	if err != nil {
		return &EmailVerificationToken{}, err
	}

	emailToken := &EmailVerificationToken{
		UserID:    userID,
		Token:     token,
		CreatedAt: time.Now(),
	}
	return emailToken, nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
