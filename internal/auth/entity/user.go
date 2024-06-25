package entity

import (
	"errors"
	"time"
)

type (
	// AuthStatus represents the status of a user's authentication.
	AuthStatus string
	// SubsStatus represents the status of a user's subscription.
	SubsStatus string
	// OAuthProvider represents the provider of the OAuth service.
	OAuthProvider string
	// UserGender represents the gender of a user.
	UserGender string
)

const (
	AuthActive   AuthStatus = "active"   // AuthActive is the active status of a user's authentication.
	AuthInactive AuthStatus = "inactive" // AuthInactive is the inactive status of a user's authentication.
	AuthLocked   AuthStatus = "locked"   // AuthLocked is the locked status of a user's authentication.

	SubsActive    SubsStatus = "active"    // SubsActive is the active status of a user's subscription.
	SubsInactive  SubsStatus = "inactive"  // SubsInactive is the inactive status of a user's subscription.
	SubsSuspended SubsStatus = "suspended" // SubsSuspended is the suspended status of a user's subscription.

	GoogleProvider   OAuthProvider = "google"   // GoogleProvider is the Google OAuth provider.
	GithubProvider   OAuthProvider = "github"   // GithubProvider is the Github OAuth provider.
	FacebookProvider OAuthProvider = "facebook" // FacebookProvider is the Facebook OAuth provider.

	MaleGender      UserGender = "male"      // MaleGender -
	FemaleGender    UserGender = "female"    // FemaleGender -.
	NonBinaryGender UserGender = "nonbinary" // NonBinaryGender -.
	OtherGender     UserGender = "other"     // OtherGender -.
)

// User represents a user entity.
type User struct {
	ID         int
	Username   string
	Email      string
	Password   string
	FullName   string
	Birthdate  time.Time
	Gender     UserGender
	Phone      string
	Status     AuthStatus
	SubsStatus SubsStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewUser creates a new instance of User.
func NewUser(username, email, password string, options ...UserOption) (*User, error) {
	user := &User{
		Username:   username,
		Email:      email,
		Password:   password,
		SubsStatus: SubsInactive,
		Status:     AuthInactive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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

func isValidStatus(status AuthStatus) bool {
	switch status {
	case AuthActive, AuthInactive, AuthLocked:
		return true
	}
	return false
}

func isValidSubsStatus(status SubsStatus) bool {
	switch status {
	case SubsActive, SubsInactive, SubsSuspended:
		return true
	}
	return false
}

func isValidGender(gender UserGender) bool {
	switch gender {
	case MaleGender, FemaleGender, NonBinaryGender, OtherGender:
		return true
	}
	return false
}

func (u *User) validate() error {
	validations := map[string]func() bool{
		"auth status":         func() bool { return isValidStatus(u.Status) },
		"subscription status": func() bool { return isValidSubsStatus(u.SubsStatus) },
		"gender":              func() bool { return isValidGender(u.Gender) },
	}

	for field, validation := range validations {
		if !validation() {
			return errors.New("invalid " + field)
		}
	}
	return nil
}
