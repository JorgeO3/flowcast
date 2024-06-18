package entity

import (
	"errors"
	"time"
)

type (
	// AuthStatus represents the status of a user's authentication.
	AuthStatus string
	// SubscriptionStatus represents the status of a user's subscription.
	SubscriptionStatus string
	// OauthProvider represents the provider of the OAuth service.
	OauthProvider string
	// UserGender represents the gender of a user.
	UserGender string
)

const (
	authActive   AuthStatus = "active"
	authInactive AuthStatus = "inactive"
	authLocked   AuthStatus = "locked"

	subscriptionActive    SubscriptionStatus = "active"
	subscriptionInactive  SubscriptionStatus = "inactive"
	subscriptionSuspended SubscriptionStatus = "suspended"

	googleProvider   OauthProvider = "google"
	githubProvider   OauthProvider = "github"
	facebookProvider OauthProvider = "facebook"

	maleGender      UserGender = "male"
	femaleGender    UserGender = "female"
	nonBinaryGender UserGender = "nonbinary"
	otherGender     UserGender = "other"
)

// User represents a user entity.
type User struct {
	ID                 int
	Username           string
	Email              string
	Password           string
	FullName           string
	Birthdate          time.Time
	Gender             UserGender
	Phone              string
	AuthStatus         AuthStatus
	SubscriptionStatus SubscriptionStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// NewUser creates a new instance of User.
func NewUser(username, email, password string, options ...UserOption) (*User, error) {
	user := &User{
		Username:           username,
		Email:              email,
		Password:           password,
		SubscriptionStatus: subscriptionInactive,
		AuthStatus:         authInactive,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	for _, option := range options {
		option(user)
	}

	// validate the enums values
	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// UserOption is a custom data type used to add more parameters to the User entity instance.
type UserOption func(*User)

// WithFullName adds the full name to the User entity instance
func WithFullName(fullName string) UserOption {
	return func(u *User) {
		u.FullName = fullName
	}
}

// WithBirthdate adds the birthdate to the User entity instance
func WithBirthdate(birthdate time.Time) UserOption {
	return func(u *User) {
		u.Birthdate = birthdate
	}
}

// WithGender adds the gender to the User entity instance
func WithGender(gender string) UserOption {
	return func(u *User) {
		u.Gender = UserGender(gender)
	}
}

// WithPhone adds the phone number to the User entity instance
func WithPhone(phone string) UserOption {
	return func(u *User) {
		u.Phone = phone
	}
}

func isValidAuthStatus(status AuthStatus) bool {
	switch status {
	case authActive, authInactive, authLocked:
		return true
	}
	return false
}

func isValidSubscriptionStatus(status SubscriptionStatus) bool {
	switch status {
	case subscriptionActive, subscriptionInactive, subscriptionSuspended:
		return true
	}
	return false
}

func isValidGender(gender UserGender) bool {
	switch gender {
	case maleGender, femaleGender, nonBinaryGender, otherGender:
		return true
	}
	return false
}

func (u *User) validate() error {
	validations := map[string]func() bool{
		"auth status":         func() bool { return isValidAuthStatus(u.AuthStatus) },
		"subscription status": func() bool { return isValidSubscriptionStatus(u.SubscriptionStatus) },
		"gender":              func() bool { return isValidGender(u.Gender) },
	}

	for field, validation := range validations {
		if !validation() {
			return errors.New("invalid " + field)
		}
	}
	return nil
}
